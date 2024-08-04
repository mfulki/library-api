package repository

import (
	"context"
	"database/sql"
	"errors"
	"user-service/internal/apperror"
	"user-service/internal/database/transaction"
	"user-service/internal/entity"

	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user entity.User) error
	SelectOneByEmail(ctx context.Context, user entity.User) (*entity.User, error)
}

type userRepositoryImpl struct {
	db transaction.Transaction
}

func NewUserRepository(db transaction.Transaction) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) RegisterUser(ctx context.Context, user entity.User) error {
	q := `
		INSERT INTO users 
			(user_name, email, user_password, date_of_birth, gender, address,photo_url)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
	`
	result, err := r.db.ExecContext(ctx, q, user.Name, user.Email, user.Password, user.DateOfBirth, user.Gender, user.Address, user.PhotoURL)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return apperror.ErrResourceNotFound
	}

	return nil
}

func (r *userRepositoryImpl) SelectOneByEmail(ctx context.Context, user entity.User) (*entity.User, error) {
	q := `SELECT user_id,user_password from users where email=$1 and deleted_at is null`
	err := r.db.QueryRowContext(ctx, q, user.Email).Scan(
		&user.ID,
		&user.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrResourceNotFound
		}

		logrus.Error(err)
		return nil, err
	}
	return &user, nil
}
