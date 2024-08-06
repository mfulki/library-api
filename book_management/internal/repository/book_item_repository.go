package repository

import (
	"book-service/internal/apperror"
	"book-service/internal/database/transaction"
	"context"

	"github.com/sirupsen/logrus"
)

type BookItemRepository interface {
	UpdateStatusBookItems(ctx context.Context, futureStatus string,currentStatus string, ids []uint64) error
}

type bookItemRepository struct {
	db transaction.Transaction
}

func NewBookItemRepository(db transaction.Transaction) *bookItemRepository {
	return &bookItemRepository{
		db: db,
	}
}

func (r *bookItemRepository) UpdateStatusBookItems(ctx context.Context, futureStatus string,currentStatus string, ids []uint64) error {
	q := `	update book_items 
			set status =$1,
				updated_at =clock_timestamp()
			where book_item_id =$2 and deleted_at is null and status=$3;`
	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		logrus.Error(err)
		return err
	}
	for _, id := range ids {
		result, err := stmt.ExecContext(ctx, futureStatus, id,currentStatus)
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
