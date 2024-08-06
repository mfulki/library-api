package server

import (
	"database/sql"
	"user-service/cmd/app/client"
	"user-service/internal/database/transaction"
	"user-service/internal/domain/custom"
	"user-service/internal/domain/example"
	"user-service/internal/handler"
	"user-service/internal/middleware"
	pbAuthor "user-service/internal/pb/author"
	pb "user-service/internal/pb/books"
	"user-service/internal/repository"
	"user-service/internal/usecase"
	"user-service/pkg/llog"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	db            transaction.Transaction
	transactor    transaction.Transactor
	fileLogger    llog.Logger
	bookService   pb.BookServiceClient
	authorService pbAuthor.AuthorServiceClient
}

func New(db *sql.DB, fileLogger *llog.FileLogger, grpcConnection *client.GRPCClient) *server {
	return &server{
		db:            transaction.NewTransaction(db),
		transactor:    transaction.NewTransactor(db),
		fileLogger:    fileLogger,
		bookService:   pb.NewBookServiceClient(grpcConnection.BookService),
		authorService: pbAuthor.NewAuthorServiceClient(grpcConnection.AuthorService),
	}
}

func (s server) Setup() *fiber.App {
	middleware := middleware.New(s.fileLogger)
	customHandler := custom.NewHandler()

	exampleRepo := example.NewRepository(s.db)
	exampleUsecase := example.NewUsecase(exampleRepo)
	exampleHandler := example.NewHandler(exampleUsecase)
	userRepository := repository.NewUserRepository(s.db)
	authUsecase := usecase.NewAuthUsecase(userRepository)
	authHandler := handler.NewAuthHandler(authUsecase)
	authorHandle := handler.NewAuthorHandler(s.authorService)
	bookHandler := handler.NewBookHandler(s.bookService, s.authorService)

	return InitRouter(&handlers{
		ExampleHandler: exampleHandler,
		AuthHandler:    authHandler,
		CustomHandler:  customHandler,
		Middleware:     middleware,
		BookHandler:    bookHandler,
		AuthorHandler:  authorHandle,
	})
}
