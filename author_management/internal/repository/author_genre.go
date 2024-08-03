package repository

import (
	"context"
	"database/sql"
	"errors"
	"library-api/author_management/internal/apperror"
	"library-api/author_management/internal/database/transaction"
	"library-api/author_management/internal/entity"

	"github.com/sirupsen/logrus"
)

type AuthorGenreRepository interface {
	SelectAllAuthorGenre(ctx context.Context) ([]entity.AuthorGenre, error)
	SelectOneAuthorGenre(ctx context.Context, authorGenre entity.AuthorGenre) (*entity.AuthorGenre, error)
	InsertAuthorGenre(ctx context.Context, authorGenres []entity.AuthorGenre) error
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

func (r *authorGenreRepositoryImpl) SelectOneAuthorGenre(ctx context.Context, authorGenre entity.AuthorGenre) (*entity.AuthorGenre, error) {
	q := `
		SELECT author_genre_id,genre_name 
		FROM author_genres
		WHERE author_genre_id=$1 and deleted_at is null
		`
	var scan entity.AuthorGenre
	if err := r.db.QueryRowContext(ctx, q, authorGenre.Id).Scan(
		&scan.Id,
		&scan.Name,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrResourceNotFound
		}

		logrus.Error(err)
		return nil, err
	}
	return &scan, nil
}

func (r *authorGenreRepositoryImpl) InsertAuthorGenre(ctx context.Context, authorGenres []entity.AuthorGenre) error {
	q := `insert into authors
		(author_name,gender,photo_url,author_genre_id) 
		VALUES
		($1,$2,$3)
	`
	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer stmt.Close()
	for _, authorGenre := range authorGenres {
		result, err := stmt.ExecContext(ctx, authorGenre.Name)
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
