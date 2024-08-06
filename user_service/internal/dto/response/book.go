package response

import (
	"user-service/internal/pb/author"
	pb "user-service/internal/pb/books"
)

type BooksResponse struct {
	Id          uint64                `json:"id,omitempty"`
	Title       string                `json:"title,omitempty"`
	Isbn        string                `json:"isbn,omitempty"`
	Description string                `json:"description,omitempty"`
	Authors     []*author.AuthorsBook `json:"author",omitempty`
	BookItem    []*pb.BookItem        `json:"book_item,omitempty"`
	Stock       uint64                `json:"stock_available,omitempty"`
	CreatedAt   string                `json:"created_at,omitempty"`
	UpdatedAt   string                `json:"updated_at,omitempty"`
	DeletedAt   string                `json:"deleted_at,omitempty"`
}

func GetBookResp(books *pb.Books, authorMap *author.AuthorsBooksMap) []BooksResponse {
	var response []BooksResponse
	for _, book := range books.Book {
		resp := BooksResponse{
			Id:          book.Id,
			Title:       book.Title,
			Isbn:        book.Isbn,
			Description: book.Description,
			BookItem:    book.BookItem,
			Authors:     authorMap.AuthorBooksMap[book.Id].AuthorBookList,
			Stock:       book.Stock,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
			DeletedAt:   book.DeletedAt,
		}
		response = append(response, resp)
	}
	return response
}
