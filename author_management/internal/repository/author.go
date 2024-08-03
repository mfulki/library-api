package repository

import (
	"context"
	"library-api/author_management/internal/apperror"
	"library-api/author_management/internal/database/transaction"
	"library-api/author_management/internal/entity"

	"github.com/sirupsen/logrus"
)

type AuthorRepository interface {
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
