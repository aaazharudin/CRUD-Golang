package main

import (
	"CRUD-Project1/config"
	"CRUD-Project1/controllers/categorycontroller"
	"CRUD-Project1/controllers/homecontroller"
	"CRUD-Project1/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	/*   1.Homepage */
	http.HandleFunc("/", homecontroller.Welcome)

	/* 2.categories */
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	/* 3.Product */
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server runing on port 8080")
	http.ListenAndServe(":8080", nil)

}
