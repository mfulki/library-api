package repository

import "user-service/internal/database/transaction"

type UserRepository interface {
}

type userRepositoryImpl struct {
	db transaction.Transaction
}

func NewUserRepository(db transaction.Transaction) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) SelectOneByEmail()
