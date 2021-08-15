package service

import (
	"github.com/ariefrpm/movies2/pkg/moviedetail"
	"github.com/ariefrpm/movies2/pkg/repository"
)

type Services struct {
	MovieDetailService moviedetail.Service
}

func CreateServices(r *repository.Repository) *Services {
	m := moviedetail.NewMovieDetailService(r.MovieRepo, r.DummyRepo)

	return &Services{
		MovieDetailService: m,
	}
}
