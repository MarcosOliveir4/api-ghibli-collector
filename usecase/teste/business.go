package teste

import (
	"api-ghibli-collector/database"
	domain "api-ghibli-collector/domain/teste"
	"api-ghibli-collector/infraestructure/repository/teste"
	"context"
)

func Listar(ctx context.Context) []domain.Teste {

	tx, errT := database.NewTransaction(ctx, true)
	if errT != nil {
		return nil
	}
	defer tx.Rollback()
	repo := teste.NewRepo(tx)

	testes, errRL := repo.List()
	if errRL != nil {
		return nil
	}

	return testes
}

func Criar(ctx context.Context, t domain.Teste) (test domain.Teste, err error) {

	tx, errT := database.NewTransaction(ctx, false)
	if errT != nil {
		return domain.Teste{}, errT
	}
	defer tx.Rollback()
	repo := teste.NewRepo(tx)

	test, err = repo.Create(t)
	if err != nil {
		return domain.Teste{}, err
	}

	err = tx.Commit()
	if err != nil {
		return domain.Teste{}, err
	}
	return test, nil
}

func BuscaPorID(ctx context.Context, id int) (test domain.Teste, err error) {

	tx, errT := database.NewTransaction(ctx, true)
	if errT != nil {
		return domain.Teste{}, errT
	}
	defer tx.Rollback()
	repo := teste.NewRepo(tx)

	test, err = repo.BuscaPorID(id)
	if err != nil {
		return domain.Teste{}, err
	}

	return test, nil
}
