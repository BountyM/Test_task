package repository

type GetResults interface {
	GetResults(data []int64) ([]int64, error)
}

type Repository struct {
	GetResults
}

func NewRepository(cache *Cache) *Repository {
	return &Repository{
		GetResults: NewGetResultsCache(cache),
	}
}
