package mysql

import (
	"database/sql"
	"log"

	"github.com/ariefrpm/movies2/pkg/library/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Saldo interface {
	SelectSaldo(accountID int64) (*SaldoDetail, error)
}

type saldo struct {
	db *sql.DB
}

func NewSaldo(db *mysql.MySql) Saldo {
	return &saldo{
		db: db.DB,
	}
}

func (s *saldo) SelectSaldo(accountID int64) (*SaldoDetail, error) {
	rows, err := s.db.Query(`select (untukKas - dariKas) saldoAkhir from(
		select sum(dariKas) dariKas, sum(untukKas) untukKas from(
			SELECT sum(nominal) darikas, 0 untukKas FROM r_transaksi_kas where dari_kas_id=?
			union all 
			SELECT 0 darikas, sum(nominal) untukKas FROM r_transaksi_kas where untuk_kas_id=?
			)x )y;`, accountID, accountID)
	if err != nil {
		log.Print("error Query", err)
	}
	var result SaldoDetail
	for rows.Next() {
		err = rows.Scan(&result.Saldo)
	}
	return &result, nil
}
