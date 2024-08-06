package repository

import (
	"category-service/internal/database/transaction"
)

type CategoryRepository interface {
}
type categoryRepositoryImpl struct {
	db transaction.Transaction
}

func NewCategoryRepository(db transaction.Transaction) *categoryRepositoryImpl {
	return &categoryRepositoryImpl{
		db: db,
	}
}
