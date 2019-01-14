package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

type (
	// DB struct
	DB struct {
		session    *mgo.Session
		collection *mgo.Collection
	}
	// Post struct
	Post struct {
		ID      bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Title   string        `bson:"title" json:"title"`
		Content string        `bson:"content" json:"content"`
	}
)

func (db *DB) getAllPost(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	w.WriteHeader(http.StatusOK)
	err := db.collection.Find(nil).All(&posts)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(posts)
		w.Write(response)
	}
}

func (db *DB) createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &post)

	// create Hash ID for new post
	post.ID = bson.NewObjectId()
	err := db.collection.Insert(post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(post)
		w.Write(response)
	}
}
func (db *DB) updatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var post Post
	putBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(putBody, &post)
	// create a Hash ID
	err := db.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": &post})
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Update successfully"))
	}
}
func (db *DB) deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := db.collection.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])})
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Deleted sucessfully"))
	}
}

func (db *DB) findPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var post Post

	w.WriteHeader(http.StatusOK)
	err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(post)
		w.Write(response)
	}
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	collection := session.DB("post").C("posts")
	db := &DB{session: session, collection: collection}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[Connected] to MongoDB")

	defer session.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", db.getAllPost).Methods("GET")
	r.HandleFunc("/api/posts", db.createPost).Methods("POST")
	r.HandleFunc("/api/posts", db.updatePost).Methods("PUT")
	r.HandleFunc("/api/posts", db.deletePost).Methods("DELETE")
	r.HandleFunc("/api/posts/{id}", db.findPost).Methods("GET")

	fmt.Println("Server running on localhost:3000")
	srv := &http.Server{
		Handler:      r,
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
