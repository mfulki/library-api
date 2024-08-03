package database

import (
	"database/sql"
	"fmt"
	"library-api/author_management/config/env"
	"library-api/author_management/internal/apperror"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgres(dbConfig *env.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	))

	if err != nil {
		return nil, apperror.Wrap(err)
	}

	if err = db.Ping(); err != nil {
		return nil, apperror.Wrap(err)
	}

	return db, nil
}
