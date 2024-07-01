package database

import (
	"context"
	"database/sql"
)

type ConnectionsPool interface {
	NewTransaction(ctx context.Context, readOnly bool) (tx Transaction, err error)
}

type Transaction interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Commit() error
	Rollback() error
}
