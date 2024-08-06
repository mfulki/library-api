package server

import (
	"book-service/internal/database/transaction"
	"book-service/internal/handler"
	"book-service/internal/repository"
	"book-service/internal/usecase"
	"book-service/pkg/llog"
	"context"
	"database/sql"
	"errors"
	"fmt"

	pb "book-service/internal/pb/books"

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
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	bookHandler := handler.NewBookHandler(bookUsecase)
	pb.RegisterBookServiceServer(s.grpcServer, bookHandler)
	return s.grpcServer
}

var noAuthMethods = map[string]bool{
	"/book.BookService/GetBook": true,
}

func jwtAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println(noAuthMethods[info.FullMethod])
	fmt.Println("Intercepting method:", info.FullMethod)
	if !noAuthMethods[info.FullMethod] {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	token := md["authorization"]

	if len(token) == 0 {
		return nil, errors.New("missing token")
	}

	return handler(ctx, req)
}
