package teste

type RepoTestes interface {
	List() ([]Teste, error)
	Create(t Teste) (Teste, error)
	BuscaPorID(id int) (Teste, error)
}
