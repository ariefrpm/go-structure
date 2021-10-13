package saldodetail

import (
	"log"

	"github.com/ariefrpm/movies2/pkg/repository/db"
)

type Service interface {
	SaldoDetail(kasId int64) (*SaldoTransaksiKasDetail, error)
}

type service struct {
	db mysql.ConnectMysql
}

func NewSaldoDetailService(db mysql.ConnectMysql) Service {
	var svc Service
	svc = &service{
		db: db,
	}
	svc = NewLoggingMiddleware(svc)
	log.Print("sini saldo 1")
	return svc
}

func (s *service) SaldoDetail(kasId int64) (*SaldoTransaksiKasDetail, error) {
	database, err := s.db.InitDB()
	
	log.Print("lanjut response")
	// var total float64 = 1000
	// var database *gorm.DB
	type SaldoAkhir struct {
		Saldo float64 `gorm:"column:saldoAkhir"`
	}
	query := `select (untukKas - dariKas) saldoAkhir from(
	select sum(dariKas) dariKas, sum(untukKas) untukKas from(
		SELECT sum(nominal) darikas, 0 untukKas FROM r_transaksi_kas where dari_kas_id=?
		union all 
		SELECT 0 darikas, sum(nominal) untukKas FROM r_transaksi_kas where untuk_kas_id=?
		)x )y;`

	var result []*SaldoAkhir
	err = database.Raw(query, kasId, kasId).Scan(&result).Error
	log.Print(result[0].Saldo)
	if err != nil {
		log.Print(err.Error())
	}
	return &SaldoTransaksiKasDetail{
		TotalSaldo: result[0].Saldo,
	}, nil
}
