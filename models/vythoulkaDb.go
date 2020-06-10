package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/go-sql-driver/mysql"
)

var VDB *gorm.DB

func NewGormDb(connectionString string) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	VDB = db
}
