package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var VDB *sql.DB

func NewDB(connectionString string) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	VDB = db
}
