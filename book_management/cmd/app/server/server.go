package server

import (
	"book-service/internal/apperror"
	"book-service/internal/constant"
	"book-service/internal/database/transaction"
	"book-service/internal/handler"
	"book-service/internal/repository"
	"book-service/internal/usecase"
	"book-service/pkg/llog"
	"book-service/pkg/utils"
	"context"
	"database/sql"
	"errors"
	"strings"

	pb "book-service/internal/pb/books"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	db         transaction.Transaction
	transactor transaction.Transactor
	fileLogger llog.Logger
	grpcServer *grpc.Server
}

func New(db *sql.DB, fileLogger *llog.FileLogger, grpcServer *grpc.Server) *server {
	return &server{
		db:         transaction.NewTransaction(db),
		transactor: transaction.NewTransactor(db),
		fileLogger: fileLogger,
		grpcServer: grpc.NewServer(grpc.UnaryInterceptor(jwtAuthInterceptor)),
	}
}

func (s server) Setup() *grpc.Server {
	bookRepository := repository.NewBookRepository(s.db)
	bookItemRepository := repository.NewBookItemRepository(s.db)
	stockJournalRepository := repository.NewStockJournalRepository(s.db)

	bookUsecase := usecase.NewBookUsecase(bookRepository, bookItemRepository, s.transactor, stockJournalRepository)
	bookHandler := handler.NewBookHandler(bookUsecase)
	pb.RegisterBookServiceServer(s.grpcServer, bookHandler)
	return s.grpcServer
}

func jwtAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var noAuthMethods = map[string]bool{
		"/book.BookService/GetBook":  true,
		"/book.BookService/GetBooks": true,
	}
	var userAuthMethods = map[string]bool{
		"/book.BookService/PostBorrows":  true,
		"/book.BookService/PostReturns": true,
	}

	if noAuthMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	if userAuthMethods[info.FullMethod] {
		return UserAuth(ctx, req, handler, md)
	}

	return handler(ctx, req)
}

func Authentication(jwtFunc func(signed string) (jwt.MapClaims, bool), ctx context.Context, md metadata.MD, req interface{}, handler grpc.UnaryHandler) (any, error) {
	authorization := md["authorization"]
	bearerToken := strings.Split(authorization[0], " ")

	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return nil, apperror.ErrUnauthorized
	}

	user, ok := jwtFunc(bearerToken[1])
	if !ok {
		return nil, apperror.ErrUnauthorized
	}

	userDataMap, ok := user["data"].(map[string]any)
	if !ok {
		return nil, apperror.ErrUnauthorized
	}

	if role, ok := user["role"]; ok {
		userDataMap["role"] = role
	}
	con := context.WithValue(ctx, constant.UserContext, userDataMap)

	return handler(con, req)
}
func UserAuth(ctx context.Context, req any, handler grpc.UnaryHandler, md metadata.MD) (any, error) {
	return Authentication(func(signed string) (jwt.MapClaims, bool) {
		user, ok := utils.JwtParseUser(signed)
		if !ok {
			return nil, false
		}
		user["role"] = "user"
		return user, true
	}, ctx, md, req, handler)
}
