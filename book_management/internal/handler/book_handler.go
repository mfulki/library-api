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
func (h *BookHandler) GetBook(ctx context.Context, in *pb.Id) (*pb.Book, error) {
	book, err := h.bookUsecase.GetBook(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return response.NewGetBookResponse(book), nil
}
func (h *BookHandler) PostBorrows(ctx context.Context, in *pb.Ids) (*pb.Message, error) {
	message, err := h.bookUsecase.UserBorrowBook(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Message{Message: *message}, nil
}
func (h *BookHandler) PostReturns(ctx context.Context, in *pb.Ids) (*pb.Message, error) {
	message, err := h.bookUsecase.UserReturnsBook(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Message{Message: *message}, nil
}
