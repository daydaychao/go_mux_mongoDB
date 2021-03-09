package main

import (
	"log"
	"net/http"

	"RESTful/go_mux_mongoDB/controllers"
	"RESTful/go_mux_mongoDB/db"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID	string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

// Author Struct
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
}


func main() {

	// Database
	db.InitDatabase()

	// API Router
	api := mux.NewRouter()
	api.HandleFunc("/api/books",controllers.GetBooks).Methods("GET")
	// api.HandleFunc("/api/books/{id}",controllers.NewABook).Methods("POST")
	// api.HandleFunc("/api/books/{id}",controllers.UpdateABook).Methods("GET")
	// api.HandleFunc("/api/books/{id}",controllers.DeleteABook).Methods("DELETE")



	println("http server :12345")
	log.Fatal(http.ListenAndServe(":12345",api))

}
