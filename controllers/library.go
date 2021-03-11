package library

import (
	"RESTful/go_mux_mongoDB/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetBooks - API - 查詢全部書
func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 查書
	var books []Book
	cursor, err := db.LbCollection.Find(ctx,bson.M{})
	if err != nil {
			log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var book Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	// 送出
	json.NewEncoder(w).Encode(books)
	fmt.Println("GetBooks done")
}

// DeleteBooks - API - 刪除全部書
func DeleteBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}

// NewBook - API - 新增 1 本書
func NewBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var newBook Book

	// 解析json
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	// 書寫入database
	insertOneResult, err := db.LbCollection.InsertOne(ctx, newBook)
	fmt.Println("NewBook End")

	if oid, ok := insertOneResult.InsertedID.(primitive.ObjectID); ok {
		json.NewEncoder(w).Encode(oid)
	}
	fmt.Println("NewBook done")
}

// GetBook - API - 用isbn查詢特定書
func GetBook(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var vars string = mux.Vars(r)["isbn"]
	fmt.Println("vars",vars)

	// 讀取第一筆資料
	// var firstBook bson.M
	// if err := db.LbCollection.FindOne(ctx, bson.M{}).Decode(&firstBook); err != nil {
	// 		log.Fatal(err)
	// }
	// fmt.Println("讀取第一筆資料",firstBook)

	// 篩選isbn
	filterCursor, err := db.LbCollection.Find(ctx, bson.M{"isbn": vars})
		if err != nil {
				log.Fatal(err)
		}
	var booksFiltered []Book
	if err = filterCursor.All(ctx, &booksFiltered); err != nil {
			log.Fatal(err)
	}

	json.NewEncoder(w).Encode(booksFiltered)
	fmt.Println("GetBook done")
}

// UpdateBook - API - 更新書(isbn)
func UpdateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}


// DeleteBook - API - 刪除書(isbn)
func DeleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}
