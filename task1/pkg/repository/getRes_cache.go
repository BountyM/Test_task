package repository

type GetResultsCache struct {
	cache *Cache
}

func NewGetResultsCache(cache *Cache) *GetResultsCache {
	return &GetResultsCache{cache: cache}
}

func (r *GetResultsCache) GetResults(data []int64) ([]int64, error) {
	return r.cache.GetSlice(data)
}
