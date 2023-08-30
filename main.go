package main

import (
	"CRUD/config"
	"CRUD/controllers/categorycontroller"
	"CRUD/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)


	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}