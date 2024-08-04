package server

import (
	"user-service/internal/domain/custom"
	"user-service/internal/domain/example"
	"user-service/internal/middleware"

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
