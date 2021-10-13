package db

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type ConnectMysql interface {
	InitDB() (*gorm.DB, error)
}

type connect struct {
}

func NewConnectionMysql() ConnectMysql {
	return &connect{}
}

func (c *connect) InitDB() (*gorm.DB, error) {
	var db *gorm.DB
	err := godotenv.Load(".env")
	// dbUri := os.Getenv("db_uri")
	// dbPass := os.Getenv("db_pass")
	// dbUser := os.Getenv("db_user")
	dsn := "mti:UX86E3vyfvo46trh@tcp(139.180.217.246:3306)/kopbidb?charset=utf8mb4&parseTime=True&loc=Local"
	log.Print(dsn)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf(err.Error())
	}
	db.LogMode(true)

	return db, err
}
