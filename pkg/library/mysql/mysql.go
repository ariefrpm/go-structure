package mysql

import (
	"database/sql"

	"github.com/jinzhu/gorm"

	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type MySql struct {
	DB *sql.DB
}

func NewMySql() *MySql {
	s, err := dbConnection()
	if err != nil {
		panic(err.Error())
	}
	return s
}

func dbConnection() (*MySql, error) {
	var db *gorm.DB
	err := godotenv.Load(".env")
	// dbUri := os.Getenv("db_uri")
	// dbPass := os.Getenv("db_pass")
	// dbUser := os.Getenv("db_user")
	dsn := "xxx:xxxx@tcp(xxx:3306)/kopbidb?charset=utf8mb4&parseTime=True&loc=Local"
	log.Print(dsn)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	db.LogMode(true)

	return &MySql{
		DB: db.DB(),
	}, err
}
