package custom

import (
	"user-service/internal/apperror"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) NotFound(ctx *fiber.Ctx) error {
	return apperror.ErrNoRoute
}
