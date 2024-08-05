package handler

import (
	"log"

	"user-service/internal/constant"
	"user-service/internal/dto"
	pb "user-service/internal/pb/books"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BookHandler struct {
	bookService pb.BookServiceClient
}

func NewBookHandler(bookService pb.BookServiceClient) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (h *BookHandler) GetAllBook(ctx *fiber.Ctx) error {
	response, err := h.bookService.GetBooks(ctx.Context(), &pb.Empty{})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}
	log.Printf("Response: %s", response)

	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Message: constant.DataRetrievedMsg,
		Data:    response,
	})
}
