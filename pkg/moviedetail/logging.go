package moviedetail

import (
	"log"
	"time"
)

type loggingMiddleware struct {
	Service
}

func NewLoggingMiddleware(s Service) Service {
	return &loggingMiddleware{s}
}

func (l *loggingMiddleware) MovieDetail(imdbID string) (m *Movie, err error)  {
	defer func(begin time.Time) {
		log.Printf("took %v, id %s, err %v\n", time.Since(begin), imdbID, err)
	}(time.Now())
	m, err = l.Service.MovieDetail(imdbID)
	return
}
