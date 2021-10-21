package repository

import (
	client "github.com/ariefrpm/movies2/pkg/library/http-client"
	libMysql "github.com/ariefrpm/movies2/pkg/library/mysql"
	"github.com/ariefrpm/movies2/pkg/repository/api"
	"github.com/ariefrpm/movies2/pkg/repository/mysql"
	"log"
)

type Repository struct {
	MovieRepo api.MovieRepo
	DummyRepo mysql.Dummy
	SaldoRepo mysql.Saldo
}


func CreateRepository() *Repository {
	log.Print("create repo")
	return &Repository{
		MovieRepo: api.NewMovieRepo(client.NewDefaultHttpClient()),
		DummyRepo: mysql.NewDummyRepo(),
		SaldoRepo: mysql.NewSaldo(libMysql.NewMySql()),
	}
}
