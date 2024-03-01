package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/dados.db")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	createTableFilms(db)

	return db
}

func createTableFilms(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS films (
			id TEXT PRIMARY KEY,
			title TEXT,
			original_title TEXT,
			original_title_romanised TEXT,
			image TEXT,
			movie_banner TEXT,
			description TEXT,
			director TEXT,
			producer TEXT,
			release_date TEXT,
			running_time TEXT,
			rt_score TEXT,
			url TEXT
		);
	`)
	if err != nil {
		panic(err.Error())
	}
}
