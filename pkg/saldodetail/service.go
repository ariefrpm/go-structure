package saldodetail

import (
	"log"

	"github.com/ariefrpm/movies2/pkg/repository/mysql"
)

type Service interface {
	SaldoDetail(kasId int64) (*SaldoTransaksiKasDetail, error)
}

type service struct {
	db mysql.Saldo
}

func NewSaldoDetailService(db mysql.Saldo) Service {
	var svc Service
	svc = &service{
		db: db,
	}
	svc = NewLoggingMiddleware(svc)
	log.Print("sini saldo 1")
	return svc
}

func (s *service) SaldoDetail(kasId int64) (*SaldoTransaksiKasDetail, error) {
	rowData, err := s.db.SelectSaldo(kasId)

	if err != nil {
		log.Print(err.Error())
	}
	return &SaldoTransaksiKasDetail{
		TotalSaldo: rowData.Saldo,
	}, nil
}
