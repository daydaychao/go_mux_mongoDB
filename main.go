package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"

	library "RESTful/go_mux_mongoDB/controllers"
	"RESTful/go_mux_mongoDB/db"
)

func main() {

	// Database initialization
	db.InitDatabase()

	// API Router
	api := mux.NewRouter()
	api.HandleFunc("/api", Index) // Welcome
	api.HandleFunc("/api/books",library.GetBooks).Methods("GET") // 查詢全部書
	api.HandleFunc("/api/books/isbn/{isbn}",library.DeleteBooks).Methods("DELETE") // 刪除全部書
	api.HandleFunc("/api/books",library.NewBook).Methods("POST") // 新增 1 本書
	api.HandleFunc("/api/books/isbn/{isbn}",library.GetBook).Methods("GET") // 查詢書(isbn)
	api.HandleFunc("/api/books/isbn/{isbn}",library.UpdateBook).Methods("POST") // 更新書(isbn)
	api.HandleFunc("/api/books/isbn/{isbn}",library.DeleteBook).Methods("DELETE") // 刪除書(isbn)


	fmt.Println("http server :12345")
	log.Fatal(http.ListenAndServe(":12345",api))
}

// Index 首頁/
func Index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode("Welcome")
}

