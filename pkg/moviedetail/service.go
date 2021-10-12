package moviedetail

import (
	"errors"
	"log"

	"github.com/ariefrpm/movies2/pkg/repository/api"
	"github.com/ariefrpm/movies2/pkg/repository/db"
)

type Service interface {
	MovieDetail(imdbID string) (*Movie, error)
}

type service struct {
	api api.MovieRepo
	db  db.Dummy
}

func NewMovieDetailService(api api.MovieRepo, db db.Dummy) Service {
	var svc Service
	svc = &service{
		api: api,
		db:  db,
	}
	svc = NewLoggingMiddleware(svc)
	log.Print("sini 1")
	return svc
}

func (s *service) MovieDetail(omdbID string) (*Movie, error) {
	log.Print("sini 2")
	if omdbID == "" {
		return nil, errors.New("imdbID is empty")
	}
	movie, err := s.api.MovieDetail(omdbID)
	if err != nil {
		return nil, err
	}

	s.db.GetSomething()

	return &Movie{
		Title:    movie.Title,
		Year:     movie.Year,
		Rated:    movie.Rated,
		Released: movie.Released,
		Runtime:  movie.Runtime,
		Genre:    movie.Genre,
		Director: movie.Director,
		Writer:   movie.Writer,
		Actors:   movie.Actors,
		Plot:     movie.Plot,
		Language: movie.Language,
		Country:  movie.Country,
		Awards:   movie.Awards,
		Poster:   movie.Poster,
		Response: movie.Response,
	}, err
}
