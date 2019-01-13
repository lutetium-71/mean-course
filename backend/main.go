package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type (
	// DB struct
	DB struct {
		client     *mongo.Client
		colleciton *mongo.Collection
	}
	// Post struct
	Post struct {
		ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		Name    string             `bson:"name" json:"name"`
		Content string             `bson:"content" json:"content"`
	}
)

func (db *DB) getAllPost(w http.ResponseWriter, r *http.Request) {
	var posts []*Post
	w.WriteHeader(http.StatusOK)
	cur, err := db.colleciton.Find(context.TODO(), nil)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		for cur.Next(context.TODO()) {
			var post Post
			err := cur.Decode(&post)
			if err != nil {
				log.Fatal(err)
			}
			posts = append(posts, &post)
		}
		if err := cur.Err(); err != nil {
			w.Write([]byte(err.Error()))
		}
		cur.Close(context.TODO())
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(posts)
		w.Write(response)
	}
	fmt.Printf("Found a single document: %+v\n", posts)
}

func (db *DB) createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &post)

	// create Hash ID for new post
	post.ID = primitive.NewObjectID()
	_, err := db.colleciton.InsertOne(context.TODO(), post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(post)
		w.Write(response)
	}

}
func (db *DB) updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}
func (db *DB) deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func (db *DB) findPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(vars["id"])
	var post Post
	filter := bson.M{"_id": id}

	w.WriteHeader(http.StatusOK)
	err := db.colleciton.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(post)
		w.Write(response)
	}
	fmt.Printf("Found a single document: %+v\n", post)
}

func main() {
	ctx, ctxErr := context.WithTimeout(context.Background(), 10*time.Second)
	if ctxErr != nil {
		log.Fatal(ctxErr)
	}
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("post").Collection("posts")
	db := &DB{client: client, colleciton: collection}

	fmt.Println("[Connected] to MongoDB")

	defer client.Disconnect(context.TODO())

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", db.getAllPost).Methods("GET")
	r.HandleFunc("/api/posts", db.createPost).Methods("POST")
	r.HandleFunc("/api/posts", db.updatePost).Methods("PUT")
	r.HandleFunc("/api/posts", db.deletePost).Methods("DELETE")
	r.HandleFunc("/api/posts/{id}", db.findPost).Methods("GET")

	fmt.Println("Server running on localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
