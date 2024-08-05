package server

import (
	"book-service/internal/database/transaction"
	"book-service/internal/handler"
	"book-service/internal/repository"
	"book-service/internal/usecase"
	"book-service/pkg/llog"
	"database/sql"

	pb "book-service/internal/pb/books"

	"google.golang.org/grpc"
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
		grpcServer: grpc.NewServer(),
	}
}

func (s server) Setup() *grpc.Server {
	bookRepository := repository.NewBookRepository(s.db)
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	bookHandler := handler.NewBookHandler(bookUsecase)
	pb.RegisterBookServiceServer(s.grpcServer, bookHandler)
	return s.grpcServer
}
