package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}
func createMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}
func updateMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}
func deleteMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}
func findMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func main() {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, "mongodb://localhost/posts")
	// if err != nil {
	// 	log.Fatal("MongoDB connection failed")
	// }
	// err = client.Ping(ctx, readpref.Primary())
	// collection := client.Database("post").Collection("posts")

	// ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	// res, err := collection.InsertOne(ctx, bson.M{"title": "Post number 3", "content": "Post content 3"})
	// id := res.InsertedID
	// fmt.Println(id)

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", getAllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/api/posts", createMoviesEndPoint).Methods("POST")
	r.HandleFunc("/api/posts", updateMoviesEndPoint).Methods("PUT")
	r.HandleFunc("/api/posts", deleteMoviesEndPoint).Methods("DELETE")
	r.HandleFunc("/api/posts/{id}", findMoviesEndPoint).Methods("GET")

	fmt.Println("Server running on localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
