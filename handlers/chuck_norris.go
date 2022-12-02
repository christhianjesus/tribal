package handlers

import (
	"errors"
	"sync"
	"tribal/entities"
	"tribal/repositories"
)

type chuckNorrisHandler struct {
	rep repositories.ChuckNorrisRepository
}

type ChuckNorrisHandler interface {
	GetManyDistinct(int) ([]entities.ApiResponse, error)
}

func NewChuckNorris(rep repositories.ChuckNorrisRepository) ChuckNorrisHandler {
	return &chuckNorrisHandler{rep}
}

func (c *chuckNorrisHandler) GetManyDistinct(quantity int) ([]entities.ApiResponse, error) {
	set := make(map[string]*entities.ApiResponse, quantity)
	missings := quantity
	tries := 0

	for missings > 0 && tries < entities.MaxTries {
		for _, v := range c.getMany(missings) {
			if v != nil {
				set[v.ID] = v
			}
		}

		missings = quantity - len(set)
		tries++
	}

	if missings != 0 {
		return nil, errors.New("maximum attempts limit exceeded")
	}

	result := make([]entities.ApiResponse, 0, len(set))
	for _, v := range set {
		result = append(result, *v)
	}

	return result, nil
}

func (c *chuckNorrisHandler) getMany(quantity int) []*entities.ApiResponse {
	var arr = make([]*entities.ApiResponse, quantity)
	var wg sync.WaitGroup

	for i := 0; i < quantity; i++ {
		wg.Add(1)

		go func(p **entities.ApiResponse) {
			defer wg.Done()

			if r, err := c.rep.Get(); err == nil {
				*p = r
			}
		}(&arr[i])
	}

	wg.Wait()

	return arr
}
