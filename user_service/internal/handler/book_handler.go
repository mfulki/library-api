package handler

import (
	"user-service/internal/constant"
	"user-service/internal/dto"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	pbAuthor "user-service/internal/pb/author"
	pb "user-service/internal/pb/books"
	pbCategory "user-service/internal/pb/categories"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BookHandler struct {
	bookService     pb.BookServiceClient
	authorService   pbAuthor.AuthorServiceClient
	categoryService pbCategory.CategoryServiceClient
}

func NewBookHandler(bookService pb.BookServiceClient,
	authorService pbAuthor.AuthorServiceClient,
	categoryService pbCategory.CategoryServiceClient,
) *BookHandler {
	return &BookHandler{
		bookService:     bookService,
		authorService:   authorService,
		categoryService: categoryService,
	}
}

func (h *BookHandler) GetAllBook(ctx *fiber.Ctx) error {
	respBook, err := h.bookService.GetBooks(ctx.Context(), &pb.Empty{})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}
	respAuthor, err := h.authorService.GetSomeAuthorsBook(ctx.Context(), &pbAuthor.Ids{Id: respBook.BookIds})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}
	respCategory, err := h.categoryService.GetSomeBookCategories(ctx.Context(), &pbCategory.Ids{Id: respBook.BookIds})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Message: constant.DataRetrievedMsg,
		Data:    response.GetBooksResp(respBook, respAuthor, respCategory),
	})
}

func (h *BookHandler) GetOneBook(ctx *fiber.Ctx) error {
	param := new(request.BookId)
	ctx.ParamsParser(&param)
	ids := []uint64{param.Id}
	respBook, err := h.bookService.GetBook(ctx.Context(), &pb.Id{Id: param.Id})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}
	respAuthor, err := h.authorService.GetSomeAuthorsBook(ctx.Context(), &pbAuthor.Ids{Id: ids})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}
	respCategory, err := h.categoryService.GetSomeBookCategories(ctx.Context(), &pbCategory.Ids{Id: ids})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Message: constant.DataRetrievedMsg,
		Data:    response.GetBookResp(respBook, respAuthor, respCategory),
	})
}
