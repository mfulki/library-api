package handler

import (
	"book-service/internal/dto/response"
	pb "book-service/internal/pb/books"
	"book-service/internal/usecase"
	"context"
)

type BookHandler struct {
	bookUsecase usecase.BookUsecase
	pb.UnimplementedBookServiceServer
}

func NewBookHandler(bookUsecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bookUsecase: bookUsecase,
	}
}
func (h *BookHandler) GetBooks(ctx context.Context, in *pb.Empty) (*pb.Books, error) {
	books, err := h.bookUsecase.GetAllBook(ctx)
	if err != nil {
		return nil, err
	}

	return response.NewGetAllBookResponse(books), nil
}
