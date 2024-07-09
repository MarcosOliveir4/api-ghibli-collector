package films

import (
	"api-ghibli-collector/database"
	domain "api-ghibli-collector/domain/films"
	"strings"
)

func NewRepo(tx database.Transaction) domain.RepoFilms {
	return &repository{tx: tx}
}

type repository struct {
	tx database.Transaction
}

func (r *repository) List(filter domain.FilmFilter) (films []domain.Film, err error) {
	var query string = "SELECT * FROM films"
	var conditions []string
	var args []interface{}
	if filter.Title != "" {
		conditions = append(conditions, "title LIKE '%' || $1 || '%'")
		args = append(args, filter.Title)
	}

	if filter.Director != "" {
		conditions = append(conditions, "director LIKE '%' || $2 || '%'")
		args = append(args, filter.Director)
	}

	if filter.RtScore != 0 {
		conditions = append(conditions, "rt_score = $3")
		args = append(args, filter.RtScore)
	}

	if filter.ReleaseDate != 0 {
		conditions = append(conditions, "release_date = $4")
		args = append(args, filter.ReleaseDate)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if filter.Limit != 0 {
		query += " LIMIT $5"
		args = append(args, filter.Limit)
	}

	if filter.Offset != 0 {
		query += " OFFSET $6"
		args = append(args, filter.Offset)
	}

	// query := `SELECT * FROM films ORDER BY release_date ASC`
	rows, err := r.tx.Query(query, args...)
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

func (r *repository) ListTotal(filter domain.FilmFilter) (total int, err error) {
	var query = `SELECT COUNT(*) FROM films`
	var conditions []string
	var args []interface{}
	if filter.Title != "" {
		conditions = append(conditions, "title LIKE '%' || $1 || '%'")
		args = append(args, filter.Title)
	}

	if filter.Director != "" {
		conditions = append(conditions, "director LIKE '%' || $2 || '%'")
		args = append(args, filter.Director)
	}

	if filter.RtScore != 0 {
		conditions = append(conditions, "rt_score = $3")
		args = append(args, filter.RtScore)
	}

	if filter.ReleaseDate != 0 {
		conditions = append(conditions, "release_date = $4")
		args = append(args, filter.ReleaseDate)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if filter.Limit != 0 {
		query += " LIMIT $5"
		args = append(args, filter.Limit)
	}

	if filter.Offset != 0 {
		query += " OFFSET $6"
		args = append(args, filter.Offset)
	}

	err = r.tx.QueryRow(query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
