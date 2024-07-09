package films

import (
	"api-ghibli-collector/database"
	domain "api-ghibli-collector/domain/films"
	"api-ghibli-collector/infraestructure/repository/films"
	"context"
)

func List(ctx context.Context, filter domain.FilmFilter) (filmsList []domain.Film, err error) {
	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	repo := films.NewRepo(tx)

	filmsList, err = repo.List(filter)
	if err != nil {
		return nil, err
	}
	return filmsList, nil
}

func ListTotal(ctx context.Context, filter domain.FilmFilter) (total int, err error) {
	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	repo := films.NewRepo(tx)

	total, err = repo.ListTotal(filter)
	if err != nil {
		return 0, err
	}
	return total, nil
}
