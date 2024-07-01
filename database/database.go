package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func (db *DB) OpenConnection() (err error) {
	dbTemp, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}

	db.db = dbTemp
	return nil
}

func (d *DB) CreatePool() (pool ConnectionsPool, err error) {
	return &connections{db: d.db}, nil
}

func (d *DB) CreateTables(pool ConnectionsPool, ctx context.Context) {
	tx, err := pool.NewTransaction(ctx, false)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS test (
        id INTEGER NOT NULL PRIMARY KEY,
        name TEXT
    );
	`
	_, err = tx.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		if err := tx.Rollback(); err != nil {
			log.Println(err)
		}
		return
	}
	_ = tx.Commit()
}

func (d *DB) CloseConnection() {
	d.db.Close()
}

func NewTransaction(ctx context.Context, readOnly bool) (tx Transaction, err error) {
	conn, ok := ctx.Value("db").(ConnectionsPool)
	if !ok {
		return nil, errors.New("database not found")
	}

	tx, err = conn.NewTransaction(ctx, readOnly)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
