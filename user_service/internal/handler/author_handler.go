package handler

import (
	"user-service/internal/constant"
	"user-service/internal/dto"
	"user-service/internal/dto/request"
	pb "user-service/internal/pb/author"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthorHandler struct {
	authorService pb.AuthorServiceClient
}

func NewAuthorHandler(authorService pb.AuthorServiceClient) *AuthorHandler {
	return &AuthorHandler{
		authorService: authorService,
	}
}

func (h *AuthorHandler) GetSomeAuthorsBook(ctx *fiber.Ctx) error {
	req := new(request.AuthorIds)
	if err := ctx.QueryParser(req); err != nil {
		return err
	}
	response, err := h.authorService.GetSomeAuthorsBook(ctx.Context(), &pb.Ids{Id: req.Ids})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Message: constant.DataRetrievedMsg,
		Data:    response,
	})
}

func (h *AuthorHandler) GetAllAuthorsBook(ctx *fiber.Ctx) error {
	req := new(request.AuthorIds)
	if err := ctx.QueryParser(req); err != nil {
		return err
	}
	response, err := h.authorService.GetAllAuthorsBook(ctx.Context(), &pb.Empty{})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Message: constant.DataRetrievedMsg,
		Data:    response,
	})
}

