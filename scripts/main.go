package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	"encoding/json"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Film struct {
	ID                     uuid.UUID `json:"id"`
	Title                  string    `json:"title"`
	OriginalTitleRomanised string    `json:"original_title_romanised"`
	Image                  string    `json:"image"`
	Description            string    `json:"description"`
	Director               string    `json:"director"`
	Producer               string    `json:"producer"`
	ReleaseDate            int64     `json:"release_date"`
	RunningTime            int64     `json:"running_time"`
	RtScore                int64     `json:"rt_score"`
}

func main() {
	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ler o arquivo json
	jsonFile, err := os.Open("../films-ghibli.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var films []Film
	err = json.Unmarshal(byteValue, &films)
	if err != nil {
		panic(err)
	}

	for _, film := range films {
		_, err := db.Exec(
			`INSERT INTO films (
				id,
				title,
				original_title_romanised,
				image,
				description,
				director,
				producer,
				release_date,
				running_time,
				rt_score
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			uuid.New().String(),
			film.Title,
			film.OriginalTitleRomanised,
			film.Image,
			film.Description,
			film.Director,
			film.Producer,
			film.ReleaseDate,
			film.RunningTime,
			film.RtScore,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("OK: ", film.Title)
	}

}
