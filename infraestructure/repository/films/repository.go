package films

import (
	"api-ghibli-collector/database"
	domain "api-ghibli-collector/domain/films"
)

func NewRepo(tx database.Transaction) domain.RepoFilms {
	return &repository{tx: tx}
}

type repository struct {
	tx database.Transaction
}

func (r *repository) List() (films []domain.Film, err error) {
	query := `SELECT * FROM films ORDER BY release_date ASC`
	rows, err := r.tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var film domain.Film
		err := rows.Scan(
			&film.ID,
			&film.Title,
			&film.OriginalTitleRomanised,
			&film.Image,
			&film.Description,
			&film.Director,
			&film.Producer,
			&film.ReleaseDate,
			&film.RunningTime,
			&film.RtScore,
		)
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}

	return films, nil
}
