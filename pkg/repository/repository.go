package repository

import (
	client "github.com/ariefrpm/movies2/pkg/library/http-client"
	"github.com/ariefrpm/movies2/pkg/repository/api"
	"github.com/ariefrpm/movies2/pkg/repository/db"
)

type Repository struct {
	MovieRepo api.MovieRepo
	DummyRepo db.Dummy
}

func CreateRepository() *Repository {
	m := api.NewMovieRepo(client.NewDefaultHttpClient())
	d := db.NewDummyRepo()

	return &Repository{
		MovieRepo: m,
		DummyRepo: d,
	}
}
