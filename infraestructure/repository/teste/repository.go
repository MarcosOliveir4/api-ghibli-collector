package teste

import (
	"api-ghibli-collector/database"
	domain "api-ghibli-collector/domain/teste"
)

func NewRepo(tx database.Transaction) domain.RepoTestes {
	return &repository{tx: tx}
}

type repository struct {
	tx database.Transaction
}

func (r *repository) List() (testes []domain.Teste, err error) {
	query := `SELECT * FROM test`
	rows, err := r.tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var teste domain.Teste
		err := rows.Scan(&teste.ID, &teste.Name)
		if err != nil {
			return nil, err
		}
		testes = append(testes, teste)
	}

	return testes, nil
}

func (r *repository) Create(t domain.Teste) (domain.Teste, error) {
	query := `INSERT INTO test (id, name) VALUES (?, ?)`
	result, err := r.tx.Exec(query, t.ID, t.Name)
	if err != nil {
		return domain.Teste{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Teste{}, err
	}

	t.ID = id
	return t, nil
}

func (r *repository) BuscaPorID(id int) (domain.Teste, error) {
	query := `SELECT * FROM test WHERE id = ?`
	row := r.tx.QueryRow(query, id)

	var teste domain.Teste
	err := row.Scan(&teste.ID, &teste.Name)
	if err != nil {
		return domain.Teste{}, err
	}

	return teste, nil
}
