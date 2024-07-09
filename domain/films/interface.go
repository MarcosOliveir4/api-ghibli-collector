package films

type RepoFilms interface {
	List(filter FilmFilter) ([]Film, error)
	ListTotal(filter FilmFilter) (int, error)
}
