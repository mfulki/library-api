package repository

import (
	"book-service/internal/apperror"
	"book-service/internal/database/transaction"
	"context"

	"github.com/sirupsen/logrus"
)

type StockJournalRepository interface {
	InsertStockJournal(ctx context.Context, futureStatus string, ids []uint64, userId uint64) error
}

type stockJournalImpl struct {
	db transaction.Transaction
}

func NewStockJournalRepository(db transaction.Transaction) *stockJournalImpl {
	return &stockJournalImpl{
		db: db,
	}
}

func (r *stockJournalImpl) InsertStockJournal(ctx context.Context, futureStatus string, ids []uint64, userId uint64) error {
	q := ` 
		Insert into stock_journals (book_item_id,status,user_id)
		Values ($1,$2,$3)
	`
	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		logrus.Error(err)
		return err
	}
	for _, id := range ids {
		result, err := stmt.ExecContext(ctx, id, futureStatus, userId)
		if err != nil {
			logrus.Error(err)
			return err
		}
		if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
			return apperror.ErrResourceNotFound
		}

	}

	return nil
}
