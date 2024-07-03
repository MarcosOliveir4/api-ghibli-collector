package films

type RepoFilms interface {
	List() ([]Film, error)
}
