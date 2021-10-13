package saldodetail

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

func (l *loggingMiddleware) SaldoDetail(kasId int64) (s *SaldoTransaksiKasDetail, err error) {
	defer func(begin time.Time) {
		log.Printf("took %v, id %s, err %v\n", time.Since(begin), kasId, err)
	}(time.Now())
	s, err = l.Service.SaldoDetail(kasId)

	return
}
