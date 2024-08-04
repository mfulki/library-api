package transaction

import (
	"context"
	"database/sql"

	"user-service/internal/constant"
)

type Transaction interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) *transaction {
	return &transaction{
		db: db,
	}
}

func (t *transaction) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if tx, ok := ctx.Value(constant.TxCtx).(*sql.Tx); ok {
		return tx.ExecContext(ctx, query, args...)
	}

	return t.db.ExecContext(ctx, query, args...)
}

func (t *transaction) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if tx, ok := ctx.Value(constant.TxCtx).(*sql.Tx); ok {
		return tx.QueryContext(ctx, query, args...)
	}

	return t.db.QueryContext(ctx, query, args...)
}

func (t *transaction) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if tx, ok := ctx.Value(constant.TxCtx).(*sql.Tx); ok {
		return tx.QueryRowContext(ctx, query, args...)
	}

	return t.db.QueryRowContext(ctx, query, args...)
}
