package mysql

import (
	"database/sql"
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
	//rows, err := s.db.Query("select saldo from tbl_saldo where acc_id = ?", accountID)
	return nil, nil
}
