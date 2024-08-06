package response

import (
	"user-service/internal/pb/author"
	pb "user-service/internal/pb/books"
	"user-service/internal/pb/categories"
)

type BooksResponse struct {
	Id          uint64                     `json:"id,omitempty"`
	Title       string                     `json:"title,omitempty"`
	Isbn        string                     `json:"isbn,omitempty"`
	Description string                     `json:"description,omitempty"`
	Authors     []*author.AuthorsBook      `json:"author,omitempty"`
	BookItem    []*pb.BookItem             `json:"book_item,omitempty"`
	Category    []*categories.BookCategory `json:"category,omitempty"`
	Stock       uint64                     `json:"stock_available,omitempty"`
	CreatedAt   string                     `json:"created_at,omitempty"`
	UpdatedAt   string                     `json:"updated_at,omitempty"`
	DeletedAt   string                     `json:"deleted_at,omitempty"`
}

func GetBookResp(books *pb.Books, authorMap *author.AuthorsBooksMap, categoryMap *categories.BookCategoriesMap) []BooksResponse {
	var response []BooksResponse
	for _, book := range books.Book {
		resp := BooksResponse{
			Id:          book.Id,
			Title:       book.Title,
			Isbn:        book.Isbn,
			Description: book.Description,
			BookItem:    book.BookItem,
			Category:    categoryMap.BookCategoriesMap[book.Id].BookCategoriesList,
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
