package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// Movie struct
type Movie struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Content string             `bson:"content" json:"content"`
}

// Trainer struct
type Trainer struct {
	Name     string `bson:"name" json:"name"`
	Age      int64  `bson:"age" json:"age"`
	Hometown string `bson:"hometown" json:"hometown"`
}

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
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[Connected] to MongoDB")

	collection := client.Database("test").Collection("trainers")

	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pweter City"}
	trainers := []interface{}{misty, brock}

	insertResult, err := collection.InsertOne(context.TODO(), ash)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	// update
	filter := bson.D{{"name", "Ash"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1}}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	// create
	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	// pass options to find method
	options := options.Find()
	options.SetLimit(2)

	var results []*Trainer

	cur, err := collection.Find(context.TODO(), nil, options)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	deleteResult, err := collection.DeleteMany(context.TODO(), nil)
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Delted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[Disconnected] from MongoDB")

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

	// r := mux.NewRouter()
	// r.HandleFunc("/api/posts", getAllMoviesEndPoint).Methods("GET")
	// r.HandleFunc("/api/posts", createMoviesEndPoint).Methods("POST")
	// r.HandleFunc("/api/posts", updateMoviesEndPoint).Methods("PUT")
	// r.HandleFunc("/api/posts", deleteMoviesEndPoint).Methods("DELETE")
	// r.HandleFunc("/api/posts/{id}", findMoviesEndPoint).Methods("GET")

	// fmt.Println("Server running on localhost:3000")
	// if err := http.ListenAndServe(":3000", r); err != nil {
	// 	log.Fatal(err)
	// }
}
