package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_fullstack")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database not connected")
	}

	log.Println("Database connected!")
}
