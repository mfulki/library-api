package server

import (
	"user-service/internal/domain/custom"
	"user-service/internal/domain/example"
	"user-service/internal/handler"
	"user-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type handlers struct {
	ExampleHandler *example.Handler
	CustomHandler  *custom.Handler
	Middleware     *middleware.Middleware
	AuthHandler    *handler.AuthHandler
	BookHandler    *handler.BookHandler
	AuthorHandler  *handler.AuthorHandler
}

func InitRouter(h *handlers) *fiber.App {
	router := fiber.New(fiber.Config{
		ErrorHandler: h.Middleware.ErrorHandler,
	})

	router.Use(recover.New())

	api := router.Group("/api")
	api.Use(h.Middleware.RequestID)
	api.Use(h.Middleware.LoggerInfo)

	auth := api.Group("/auth")
	auth.Post("/login", h.AuthHandler.Login)
	auth.Post("/register", h.AuthHandler.Register)

	book := api.Group("/book")
	book.Get("/", h.BookHandler.GetAllBook)
	book.Get("/:id", h.BookHandler.GetOneBook)

	user := book.Group("/")
	user.Use(h.Middleware.UserAuth())
	user.Post("/borrow", h.BookHandler.PostBorrow)
	user.Post("/return", h.BookHandler.PostReturn)

	author := api.Group("/author")
	author.Get("/some", h.AuthorHandler.GetSomeAuthorsBook)
	author.Get("/", h.AuthorHandler.GetAllAuthorsBook)

	h.ExampleHandler.RegisterRoute(api)

	router.Use(h.CustomHandler.NotFound)

	return router
}
