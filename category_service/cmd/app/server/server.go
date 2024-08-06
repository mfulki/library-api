package server

import (
	"category-service/internal/database/transaction"
	"category-service/internal/handler"
	"category-service/internal/repository"
	"category-service/internal/usecase"
	"category-service/pkg/llog"
	"database/sql"

	pb "category-service/internal/pb/categories"

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
	categoryRepository := repository.NewCategoryRepository(s.db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	pb.RegisterCategoryServiceServer(s.grpcServer, categoryHandler)
	return s.grpcServer
}
