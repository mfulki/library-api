package handler

import (
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
	bookItem := []*pb.BookItem{
		{Id: 1,
			BookId: 1,
			Status: "available"},
		{Id: 2,
			BookId: 1,
			Status: "available"},
	}
	books := &pb.Books{
		Book: []*pb.Book{
			{Id: 1,
				Title:       "sendy",
				Isbn:        "12-4-5-6-6-",
				Description: "mermead",
				AuthorId:    []uint64{1, 2, 3, 4, 5},
				CategoryId:  []uint64{1, 2, 3, 4, 5},
				BookItem:    bookItem,
				CreatedAt:   "12-11-11",
				UpdatedAt:   "12-11-11",
				DeletedAt:   "12-11-11",
			},
		},
	}
	return books, nil
}
