package films

import (
	"api-ghibli-collector/database"
	domain "api-ghibli-collector/domain/films"
	"api-ghibli-collector/infraestructure/repository/films"
	"context"
)

func List(ctx context.Context) (filmsList []domain.Film, err error) {
	tx, err := database.NewTransaction(ctx, true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	repo := films.NewRepo(tx)

	filmsList, err = repo.List()
	if err != nil {
		return nil, err
	}
	return filmsList, nil
}
