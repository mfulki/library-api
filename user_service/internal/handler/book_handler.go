package handler

import "github.com/gofiber/fiber/v2"

type BookHandler struct {

}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (h *BookHandler) GetAllBook(ctx *fiber.Ctx) error {
	return nil
}
