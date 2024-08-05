package repository

import (
	"book-service/internal/database/transaction"
	"book-service/internal/entity"
	"context"
	"log"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type BookRepository interface {
	GetAllBook(ctx context.Context) ([]entity.Book, error)
}
type bookRepositoryImpl struct {
	db transaction.Transaction
}

func NewBookRepository(db transaction.Transaction) *bookRepositoryImpl {
	return &bookRepositoryImpl{
		db: db,
	}
}

func (r *bookRepositoryImpl) GetAllBook(ctx context.Context) ([]entity.Book, error) {
	q := `
    SELECT 
        b.book_id,
        b.book_title,
        b.isbn,
        b.description,
        b.created_at AS book_created_at,
        b.updated_at AS book_updated_at,
        b.deleted_at AS book_deleted_at,
        ARRAY_AGG(DISTINCT bc.category_id) AS categories,
        ARRAY_AGG(DISTINCT bi.book_item_id) AS book_items,
        ARRAY_AGG(DISTINCT ba.author_id) AS authors,
		jsonb_AGG(Distinct
        jsonb_build_object(
                'book_item_id', bi.book_item_id,
                'status', bi.status))
        AS book_item_statuses
    FROM books b
    LEFT JOIN book_categories bc ON b.book_id = bc.book_id
    LEFT JOIN book_items bi ON b.book_id = bi.book_id
    LEFT JOIN book_authors ba ON b.book_id = ba.book_id
	WHERE b.deleted_at is null and bc.deleted_at is null and ba.deleted_at is null
    GROUP BY 
        b.book_id, 
        b.book_title, 
        b.isbn, 
        b.description, 
        b.created_at, 
        b.updated_at, 
        b.deleted_at,
		bi.status;
    `
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		var category pq.Int64Array
		var author pq.Int64Array
		var bookItem pq.Int64Array

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.ISBN,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.DeletedAt,
			&category,
			&bookItem,
			&author,
			&book.BookItem,
		)
		if err != nil {
			log.Fatalf("Row scanning failed: %v", err)
			return nil, err
		}
		book.CategoryId = []int64(category)
		book.AuthorId = []int64(author)
		books = append(books, book)

	}
	return books, nil
}
