package repository

import (
	"category-service/internal/database/transaction"
	"category-service/internal/entity"
	"context"

	"github.com/sirupsen/logrus"
)

type CategoryRepository interface {
	GetSomeBookCategories(ctx context.Context, ids []uint64) (map[uint64]entity.BookCategoryJson, error)
}
type categoryRepositoryImpl struct {
	db transaction.Transaction
}

func NewCategoryRepository(db transaction.Transaction) *categoryRepositoryImpl {
	return &categoryRepositoryImpl{
		db: db,
	}
}

func (r *categoryRepositoryImpl) GetSomeBookCategories(ctx context.Context, ids []uint64) (map[uint64]entity.BookCategoryJson, error) {
	q := ` select b.book_id,jsonb_AGG(Distinct
        jsonb_build_object(
                'category_id', c.category_id ,
                'category_name', c.category_name))
        AS book_item_statuses from book_categories  b 
        left join categories c on c.category_id=b.category_id
		WHERE book_id = ANY($1)
        group by b.book_id;`
	rows, err := r.db.QueryContext(ctx, q, ids)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	var categoryBooks entity.BookCategoryJson
	results := make(map[uint64]entity.BookCategoryJson, 0)
	for rows.Next() {
		var bookId uint64
		if err := rows.Scan(&bookId, &categoryBooks); err != nil {
			logrus.Error(err)
			return nil, err
		}
		results[bookId] = categoryBooks
	}

	if err := rows.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return results, nil
}
