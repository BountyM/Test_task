package repository

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	resMap map[int64]int64
}

func NewCache() *Cache {
	return &Cache{
		resMap: make(map[int64]int64),
	}
}

func (c *Cache) Set(data, res int64) {
	c.Lock()
	defer c.Unlock()
	c.resMap[data] = res
}

func (c *Cache) Get(data int64) (int64, error) {
	c.RLock()
	res, ok := c.resMap[data]
	c.RUnlock()
	if !ok {
		res, err := GetResult(data)
		if err != nil {
			return res, err
		}
		c.Set(data, res)
		return res, nil
	} else {
		return res, nil
	}
}

func (c *Cache) GetSlice(data []int64) ([]int64, error) {
	ld := len(data)
	res := make([]int64, ld)
	var wg sync.WaitGroup
	wg.Add(ld)
	for i, d := range data {
		go func(i int, d int64, wg *sync.WaitGroup) {
			defer wg.Done()
			a, err := c.Get(d)
			if err != nil {
				log.Println("error cache get :", err)
				a = 0
			}

			res[i] = a
		}(i, d, &wg)

	}
	wg.Wait()
	return res, nil
}

func GetResult(data int64) (int64, error) {
	time.Sleep(10 * time.Second)
	return data + 14, nil
}
