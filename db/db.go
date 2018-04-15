package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"zine/catalog"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *gorm.DB
)

func InitDB() {
	DBCon, _  = gorm.Open("postgres",
		"host=127.0.0.1 port=5432 user=postgres dbname=dev password=password sslmode=disable")
	DBCon.LogMode(true)

	DBCon.Debug().AutoMigrate(&catalog.Publication{}, &catalog.Magazine{}, &catalog.Asset{})
	//defer DBCon.Close()
}
