package main

import (
	"CRUD/config"
	"CRUD/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}