package example

import (
	"context"

	"book-service/internal/apperror"
	"book-service/internal/database/transaction"
)

type Repository interface {
	SelectOneByID(context.Context, int) (int, error)
	InsertOne(context.Context, Entity) (int, error)
}

type repositoryImpl struct {
	db transaction.Transaction
}

func NewRepository(db transaction.Transaction) *repositoryImpl {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) SelectOneByID(ctx context.Context, exampleID int) (int, error) {
	q := "SELECT 5"

	var scan int
	if err := r.db.QueryRowContext(ctx, q).Scan(&scan); err != nil {
		return 0, apperror.Wrap(err)
	}

	return scan + exampleID, nil
}

func (r *repositoryImpl) InsertOne(ctx context.Context, newExample Entity) (int, error) {
	q := "SELECT 10"

	var scan int
	if err := r.db.QueryRowContext(ctx, q).Scan(&scan); err != nil {
		return 0, apperror.Wrap(err)
	}

	return scan + int(newExample.ID), nil
}
