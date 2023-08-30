package main

import (
	"CRUD/config"
	"CRUD/controllers/categorycontroller"
	"CRUD/controllers/homecontroller"
	"CRUD/controllers/productcontroller"
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

	// Product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}