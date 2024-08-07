package response

import (
	"book-service/internal/entity"
	pb "book-service/internal/pb/books"
)

func NewGetAllBookResponse(books *entity.Books) *pb.Books {
	pbBooks := []*pb.Book{}
	for _, book := range books.Slice {
		bookItem := []*pb.BookItem{}
		for _, item := range book.BookItem {
			bookItem = append(bookItem, &pb.BookItem{
				Id:     item.Id,
				BookId: book.Id,
				Status: item.Status,
			})
		}
		pbBook := pb.Book{
			Id:          book.Id,
			Title:       book.Title,
			Description: book.Description,
			AuthorId:    book.AuthorId,
			CategoryId:  book.CategoryId,
			BookItem:    bookItem,
			Stock:       book.Stock,
			CreatedAt:   book.CreatedAt.Time.String(),
			UpdatedAt:   book.UpdatedAt.Time.String(),
			DeletedAt:   book.DeletedAt.Time.String(),
		}
		pbBooks = append(pbBooks, &pbBook)
	}
	return &pb.Books{Book: pbBooks, BookIds: books.BookIds}
}
func NewGetBookResponse(book *entity.Book) *pb.Book {
	bookItem := []*pb.BookItem{}
	for _, item := range book.BookItem {
		bookItem = append(bookItem, &pb.BookItem{
			Id:     item.Id,
			BookId: book.Id,
			Status: item.Status,
		})
	}
	return &pb.Book{
		Id:          book.Id,
		Title:       book.Title,
		Description: book.Description,
		AuthorId:    book.AuthorId,
		CategoryId:  book.CategoryId,
		BookItem:    bookItem,
		Stock:       book.Stock,
		CreatedAt:   book.CreatedAt.Time.String(),
		UpdatedAt:   book.UpdatedAt.Time.String(),
		DeletedAt:   book.DeletedAt.Time.String(),
	}

}
