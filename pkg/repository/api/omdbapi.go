package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	client "github.com/ariefrpm/movies2/pkg/library/http-client"
)

type MovieRepo interface {
	//SearchMovie(page int, word string) (*entity.MovieList, error)
	MovieDetail(imdbID string) (*Movie, error)
}

type movieRepo struct {
	client client.HttpClient
}

type saldoRepo struct {
	client client.HttpClient
}

func NewMovieRepo(client client.HttpClient) MovieRepo {
	return &movieRepo{
		client: client,
	}
}

const (
	BaseUrl = "http://www.omdbapi.com"
	OmdbKey = "faf7e5bb"
)

func (m *movieRepo) MovieDetail(imdbID string) (*Movie, error) {
	log.Print("sini 4")
	url := fmt.Sprintf("%s/?apikey=%s&i=%s", BaseUrl, OmdbKey, imdbID)

	data, err := m.client.GET(url)
	if err != nil {
		return nil, err
	}

	var res Movie
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	if res.Response != "True" {
		return nil, errors.New(fmt.Sprintf("response error: %s", res.Response))
	}

	return &res, nil
}
