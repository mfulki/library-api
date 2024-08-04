package server

import (
	"book-service/internal/domain/custom"
	"book-service/internal/domain/example"
	"book-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type handlers struct {
	ExampleHandler *example.Handler
	CustomHandler  *custom.Handler
	Middleware     *middleware.Middleware
}

func InitRouter(h *handlers) *fiber.App {
	router := fiber.New(fiber.Config{
		ErrorHandler: h.Middleware.ErrorHandler,
	})

	router.Use(recover.New())

	api := router.Group("/api")
	api.Use(h.Middleware.RequestID)
	api.Use(h.Middleware.LoggerInfo)

	h.ExampleHandler.RegisterRoute(api)

	api.Use(h.CustomHandler.NotFound)

	return router
}
