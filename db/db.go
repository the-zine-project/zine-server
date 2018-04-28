package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *gorm.DB
)

func InitDB() {
	DBCon, _ = gorm.Open("postgres",
		"host=127.0.0.1 port=5432 user=postgres dbname=dev password=password sslmode=disable")
	DBCon.LogMode(true)
}
