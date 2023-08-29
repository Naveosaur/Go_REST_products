package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/go_rest_api")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	DB = db
}