package database

import "database/sql"

type transaction struct {
	tx *sql.Tx
}

func (t *transaction) Commit() error {
	return t.tx.Commit()
}

func (t *transaction) Rollback() error {
	return t.tx.Rollback()
}

func (t *transaction) Exec(query string, args ...any) (sql.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t *transaction) Query(query string, args ...any) (*sql.Rows, error) {
	return t.tx.Query(query, args...)
}

func (t *transaction) QueryRow(query string, args ...any) *sql.Row {
	return t.tx.QueryRow(query, args...)
}
