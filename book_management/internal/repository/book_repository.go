package repository

import "book-service/internal/database/transaction"

type BookRepository interface {

}
type bookRepositoryImpl struct {
	db transaction.Transaction
}

func NewBookRepository(db transaction.Transaction)*bookRepositoryImpl{
	return &bookRepositoryImpl{
		db: db,
	}
}
