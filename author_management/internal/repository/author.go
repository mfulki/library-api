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

type AuthorRepository interface {
	SelectOneAuthor(ctx context.Context, author entity.Author) (*entity.Author, error)
	SelectAllAuthor(ctx context.Context) ([]entity.Author, error)
	InsertAuthor(ctx context.Context, authors []entity.Author) error
}

type authorRepositoryImpl struct {
	db transaction.Transaction
}

func NewAuthorRepository(db transaction.Transaction) *authorRepositoryImpl {
	return &authorRepositoryImpl{
		db: db,
	}
}

func (r *authorRepositoryImpl) InsertAuthor(ctx context.Context, authors []entity.Author) error {
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
	for _, author := range authors {
		result, err := stmt.ExecContext(ctx, author.Name, author.Gender, author.PhotoUrl, author.Genre.Id)
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

func (r *authorRepositoryImpl) SelectOneAuthor(ctx context.Context, author entity.Author) (*entity.Author, error) {
	q := `
		SELECT a.author_id,a.author_name,a.photo_url,a.gender,g.author_genre_id,g.genre_name 
		FROM authors a
		JOIN author_genres g ON a.author_genre_id=g.author_genre_id
		WHERE author_id=$1 and deleted_at is null
		`
	var scan entity.Author
	if err := r.db.QueryRowContext(ctx, q, author.Id).Scan(
		&scan.Id,
		&scan.Name,
		&scan.PhotoUrl,
		&scan.Genre.Id,
		&scan.Genre.Name,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrResourceNotFound
		}

		logrus.Error(err)
		return nil, err
	}
	return &scan, nil
}

func (r *authorRepositoryImpl) SelectAllAuthor(ctx context.Context) ([]entity.Author, error) {
	q := `
		SELECT a.author_id,a.author_name,a.photo_url,a.gender,g.author_genre_id,g.genre_name 
		FROM authors a
		JOIN author_genres g ON a.author_genre_id=g.author_genre_id
		WHERE deleted_at is null
		`
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	results := make([]entity.Author, 0)
	for rows.Next() {
		var scan entity.Author
		if err := rows.Scan(&scan.Id, &scan.Name, &scan.PhotoUrl, &scan.Gender, &scan.Genre.Id, &scan.Genre.Name); err != nil {
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
