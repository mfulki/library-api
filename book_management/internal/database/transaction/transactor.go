package transaction

import (
	"book-service/internal/apperror"
	"book-service/internal/constant"
	"context"
	"database/sql"
)

type Transactor interface {
	WithTransaction(context.Context, func(context.Context) (any, error)) (any, error)
}

type transactor struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) *transactor {
	return &transactor{db: db}
}

func (t *transactor) WithTransaction(ctx context.Context, tFunc func(context.Context) (any, error)) (any, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, apperror.Wrap(err)
	}

	returnedData, err := tFunc(context.WithValue(ctx, constant.TxCtx, tx))
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, apperror.Wrap(err)
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, apperror.Wrap(err)
	}

	return returnedData, nil
}
