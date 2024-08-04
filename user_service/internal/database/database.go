package database

import (
	"database/sql"
	"fmt"
	"user-service/config/env"
	"user-service/internal/apperror"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgress(dbCfg *env.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.Port,
	))

	if err != nil {
		return nil, apperror.Wrap(err)
	}

	if err = db.Ping(); err != nil {
		return nil, apperror.Wrap(err)
	}

	return db, nil
}
