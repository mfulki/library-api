package repository

import (
	"context"
	"library-api/author_management/internal/database/transaction"
	"library-api/author_management/internal/entity"

	"github.com/sirupsen/logrus"
)

type AuthorGenreRepository interface {
	SelectAllAuthorGenre(ctx context.Context) ([]entity.AuthorGenre, error)
}

type authorGenreRepositoryImpl struct {
	db transaction.Transaction
}

func NewAuthorGenreRepository(db transaction.Transaction) *authorGenreRepositoryImpl {
	return &authorGenreRepositoryImpl{
		db: db,
	}
}

func (r *authorGenreRepositoryImpl) SelectAllAuthorGenre(ctx context.Context) ([]entity.AuthorGenre, error) {
	q := `
		SELECT author_genre_id,genre_name 
		FROM author_genres
		WHERE deleted_at is null`
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	results := make([]entity.AuthorGenre, 0)
	for rows.Next() {
		var scan entity.AuthorGenre
		if err := rows.Scan(&scan.Id, &scan.Name); err != nil {
			logrus.Error(err)
			return nil, err
		}

		results = append(results, scan)
	}

	if err := rows.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return results, nil
}
