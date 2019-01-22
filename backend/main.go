package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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
	// User struct
	User struct {
		ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Email    string        `bson:"email" json:"email"`
		Password string        `bson:"password" json:"password"`
	}
)

func (db *DB) getAllPosts(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusCreated)
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

	// update post
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

func (db *DB) getPost(w http.ResponseWriter, r *http.Request) {
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

func (db *DB) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	userBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(userBody, &user)

	// create Hash ID for new user
	user.ID = bson.NewObjectId()
	user.Password = hashAndSalt(getPassword(user.Password))
	err := db.collection.Insert(user)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(user)
		w.Write(response)
	}
}

func getPassword(password string) []byte {
	return []byte(password)
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
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

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Content-Type", "Accept"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", db.getAllPosts).Methods("GET")
	r.HandleFunc("/api/posts", db.createPost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", db.updatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", db.deletePost).Methods("DELETE")
	r.HandleFunc("/api/posts/{id}", db.getPost).Methods("GET")

	r.HandleFunc("/api/user/signup", db.createUser).Methods("POST")

	fmt.Println("Server running on localhost:3000")
	http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
