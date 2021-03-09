package controllers

import (
	"RESTful/go_mux_mongoDB/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetBooks - Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var book bson.M
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := db.DbCollection.Find(ctx, bson.M{})
	if err != nil {
			log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
			log.Fatal(err)
	}
	fmt.Println(episodes)

	if err = db.DbCollection.FindOne(ctx, bson.M{}).Decode(&book); err != nil {
			log.Fatal(err)
	}
	fmt.Println(book)

	json.NewEncoder(w).Encode(book)
}
