package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

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

var signinKey = []byte("secret")

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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to Create User")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(user)
	w.Write(response)
}

func (db *DB) loginUser(w http.ResponseWriter, r *http.Request) {
	message := make(map[string]string)
	var inputUser User
	var authUser User

	// decode into User struct
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		message["message"] = "Error in request body"
		fmt.Fprintln(w, "Error in request body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(message)
		w.Write(response)
		return
	}
	// validate user credentials
	err = db.collection.Find(bson.M{"email": inputUser.Email}).One(&authUser)
	if err != nil {
		fmt.Println(inputUser)
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Invalid User")
		message["message"] = "Invalid User"
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(message)
		w.Write(response)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(inputUser.Password))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Invalid Password")
		message["message"] = "Invalid Password"
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(message)
		w.Write(response)
		return
	}
	// create a signer for rsa256
	validToken, err := generateJWT(authUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while Signing Token!")
		log.Printf("Token Signing error: %v\n", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Token", validToken)
	t := make(map[string]string)
	t["Token"] = validToken
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(t)
	w.Write(response)
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

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return signinKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func generateJWT(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
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
	r.HandleFunc("/api/user/login", db.loginUser).Methods("POST")

	fmt.Println("Server running on localhost:3000")
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
