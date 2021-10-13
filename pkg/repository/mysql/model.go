package mysql

type Some struct {
	Name string
}

type SaldoDetail struct {
	Saldo float64 `db:"saldo"`
}
