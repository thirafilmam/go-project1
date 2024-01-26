package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connectdb() {
	db, err := sql.Open("mysql", "root:1234@/data_kampus?parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("Database Berjalan")
	DB = db
}
