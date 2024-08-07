package repository

import (
	"author-service/internal/apperror"
	"author-service/internal/database/transaction"
	"author-service/internal/entity"
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

type AuthorRepository interface {
	GetOneAuthor(ctx context.Context, author entity.Author) (*entity.Author, error)
	GetAllAuthor(ctx context.Context) ([]entity.Author, error)
	InsertAuthor(ctx context.Context, authors []entity.Author) error
	GetSomeAuthor(ctx context.Context, ids []uint64) ([]entity.Author, error)
	DeleteOneAuthor(ctx context.Context, id uint64) error
	GetAllAuthorsBook(ctx context.Context) (map[uint64]entity.AuthorsBooksJson, error)
	GetSomeAuthorsBook(ctx context.Context, ids []uint64) (map[uint64]entity.AuthorsBooksJson, error)
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
		(author_name,gender,photo_url) 
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
		result, err := stmt.ExecContext(ctx, author.Name, author.Gender, author.PhotoUrl)
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

func (r *authorRepositoryImpl) GetOneAuthor(ctx context.Context, author entity.Author) (*entity.Author, error) {
	q := `
		SELECT a.author_id,a.author_name,a.photo_url,a.gender,g.author_genre_id
		FROM authors a
		WHERE author_id=$1 and deleted_at is null
		`
	var scan entity.Author
	if err := r.db.QueryRowContext(ctx, q, author.Id).Scan(
		&scan.Id,
		&scan.Name,
		&scan.PhotoUrl,
		&scan.Gender,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrResourceNotFound
		}

		logrus.Error(err)
		return nil, err
	}
	return &scan, nil
}

func (r *authorRepositoryImpl) GetAllAuthor(ctx context.Context) ([]entity.Author, error) {
	q := `
		SELECT a.author_id,a.author_name,a.photo_url,a.gender
		FROM authors a
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
		if err := rows.Scan(&scan.Id, &scan.Name, &scan.PhotoUrl, &scan.Gender); err != nil {
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

func (r *authorRepositoryImpl) GetSomeAuthor(ctx context.Context, ids []uint64) ([]entity.Author, error) {
	q := `
		SELECT a.author_id,a.author_name,a.photo_url,a.gender
		FROM authors a
		WHERE deleted_at is null and a.author_id=$1
		`
	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var authors []entity.Author
	for _, id := range ids {
		var author entity.Author
		if err := stmt.QueryRowContext(ctx, id).Scan(
			&author.Id, &author.Name, &author.PhotoUrl, &author.Gender); err != nil {
			logrus.Error(err)
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (r *authorRepositoryImpl) DeleteOneAuthor(ctx context.Context, id uint64) error {
	q := `UPDATE authors set deleted_at=now where author_id=$1`
	result, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return apperror.ErrResourceNotFound
	}

	return nil
}
func (r *authorRepositoryImpl) GetAllAuthorsBook(ctx context.Context) (map[uint64]entity.AuthorsBooksJson, error) {
	q := `select book_id,jsonb_AGG(Distinct
        jsonb_build_object(
                'author_id', a.author_id ,
                'author_name', a.author_name))
        AS book_item_statuses from author_books b 
        left join authors a on a.author_id=b.author_id 
        group by b.book_id;`
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	var authorsBooks entity.AuthorsBooksJson
	results := make(map[uint64]entity.AuthorsBooksJson, 0)
	for rows.Next() {
		var bookId uint64
		if err := rows.Scan(&bookId, &authorsBooks); err != nil {
			logrus.Error(err)
			return nil, err
		}
		results[bookId] = authorsBooks
	}

	if err := rows.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return results, nil
}

func (r *authorRepositoryImpl) GetSomeAuthorsBook(ctx context.Context, ids []uint64) (map[uint64]entity.AuthorsBooksJson, error) {
	q := `select book_id,jsonb_AGG(Distinct
        jsonb_build_object(
                'author_id', a.author_id ,
                'author_name', a.author_name))
        AS book_item_statuses from author_books b 
        left join authors a on a.author_id=b.author_id 
		WHERE book_id = ANY($1)
        group by b.book_id;`
	rows, err := r.db.QueryContext(ctx, q, ids)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	var authorsBooks entity.AuthorsBooksJson
	results := make(map[uint64]entity.AuthorsBooksJson, 0)
	for rows.Next() {
		var bookId uint64
		if err := rows.Scan(&bookId, &authorsBooks); err != nil {
			logrus.Error(err)
			return nil, err
		}
		results[bookId] = authorsBooks
	}

	if err := rows.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return results, nil
}
