package main

import (
	"go-native/config"
	"go-native/controllers/categorycontroller"
	"go-native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.Connectdb()

	//homepage call
	http.HandleFunc("/", homecontroller.Welcome)

	//CRUD
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
