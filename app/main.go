package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	userid int
	name   string
	email  string
}

func ProfilePicture(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./profile.png")
}

func GetProfile(w http.ResponseWriter, req *http.Request) {
	// mongodb part
	// create client and connect
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	// Close connection defer
	defer func() {
		err = client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed")
	}()

	collection := client.Database("user-account").Collection("users")

	user := User{}
	err = collection.FindOne(context.TODO(), bson.M{"userid": 1}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	fmt.Println("Start server 127.0.0.1:3000")
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/profile-picture", ProfilePicture)
	http.HandleFunc("/get-profile", GetProfile)

	http.ListenAndServe(":3000", nil)
}
