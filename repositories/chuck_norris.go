package repositories

import (
	"encoding/json"
	"io"
	"net/http"
	"tribal/entities"
)

type chuckNorris struct {
	url string
}

type ChuckNorrisRepository interface {
	Get() (*entities.ApiResponse, error)
}

func NewChuckNorris(url string) ChuckNorrisRepository {
	return &chuckNorris{url}
}

func (c *chuckNorris) Get() (*entities.ApiResponse, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var response entities.ApiResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
