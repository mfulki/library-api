package server

import (
	"author-service/internal/database/transaction"
	"author-service/internal/handler"
	pb "author-service/internal/pb/author"
	"author-service/internal/repository"
	"author-service/internal/usecase"
	"author-service/pkg/llog"
	"database/sql"

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
	authorRepository := repository.NewAuthorRepository(s.db)
	authorUsecase := usecase.NewAuthorUsecase(authorRepository)
	authorHandler := handler.NewAuthorHandler(authorUsecase)
	pb.RegisterAuthorServiceServer(s.grpcServer, authorHandler)
	return s.grpcServer
}
