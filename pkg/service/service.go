package service

import (
	"log"

	"github.com/ariefrpm/movies2/pkg/moviedetail"
	"github.com/ariefrpm/movies2/pkg/repository"
	"github.com/ariefrpm/movies2/pkg/saldodetail"
)

type Services struct {
	MovieDetailService moviedetail.Service
	SaldoDetailService saldodetail.Service
}

func CreateServices(r *repository.Repository) *Services {
	log.Print("kesini yaa")
	m := moviedetail.NewMovieDetailService(r.MovieRepo, r.DummyRepo)
	s := saldodetail.NewSaldoDetailService(r.SaldoRepo)
	return &Services{
		MovieDetailService: m,
		SaldoDetailService: s,
	}
}
