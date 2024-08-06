package handler

import (
	"user-service/internal/constant"
	"user-service/internal/dto"
	pbAuthor "user-service/internal/pb/author"
	pb "user-service/internal/pb/books"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BookHandler struct {
	bookService   pb.BookServiceClient
	authorService pbAuthor.AuthorServiceClient
}

func NewBookHandler(bookService pb.BookServiceClient, authorService pbAuthor.AuthorServiceClient) *BookHandler {
	return &BookHandler{
		bookService:   bookService,
		authorService: authorService,
	}
}

func (h *BookHandler) GetAllBook(ctx *fiber.Ctx) error {
	response, err := h.bookService.GetBooks(ctx.Context(), &pb.Empty{})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}


	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Message: constant.DataRetrievedMsg,
		Data:    response,
	})
}