package repository

import (
	"log"

	client "github.com/ariefrpm/movies2/pkg/library/http-client"
	"github.com/ariefrpm/movies2/pkg/repository/api"
	"github.com/ariefrpm/movies2/pkg/repository/db"
)

type Repository struct {
	MovieRepo api.MovieRepo
	DummyRepo db.Dummy
	DBMysql   db.ConnectMysql
}

func CreateRepository() *Repository {
	m := api.NewMovieRepo(client.NewDefaultHttpClient())
	d := db.NewDummyRepo()
	dm := db.NewConnectionMysql()
	log.Print("create repo")
	return &Repository{
		MovieRepo: m,
		DummyRepo: d,
		DBMysql:   dm,
	}
}
