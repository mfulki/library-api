package server

import (
	"database/sql"
	"user-service/internal/database/transaction"
	"user-service/internal/domain/custom"
	"user-service/internal/domain/example"
	"user-service/internal/middleware"
	"user-service/pkg/llog"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	db         transaction.Transaction
	transactor transaction.Transactor
	fileLogger llog.Logger
}

func New(db *sql.DB, fileLogger *llog.FileLogger) *server {
	return &server{
		db:         transaction.NewTransaction(db),
		transactor: transaction.NewTransactor(db),
		fileLogger: fileLogger,
	}
}

func (s server) Setup() *fiber.App {
	middleware := middleware.New(s.fileLogger)
	customHandler := custom.NewHandler()

	exampleRepo := example.NewRepository(s.db)
	exampleUsecase := example.NewUsecase(exampleRepo)
	exampleHandler := example.NewHandler(exampleUsecase)

	return InitRouter(&handlers{
		ExampleHandler: exampleHandler,
		CustomHandler:  customHandler,
		Middleware:     middleware,
	})
}
