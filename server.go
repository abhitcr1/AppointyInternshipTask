package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title,omitempty"`
	Body  string             `json:"body" bson:"body,omitempty"`
	Tags  string             `json:"tags" bson:"tags,omitempty"`
}
var collection = ConnecttoDB()
func main() {
	//Init Router
	router := httprouter.New()

	//Routing for different HTTP methods
	router.GET("/", showHome)
	router.GET("/article", getUserPosts)
	router.GET("/article/:id", getUser)
	router.GET("/articles/search?q=title", searchUser)
	router.POST("/articles", createPost)
	router.POST("/articles", createUser)
	// set our port address as 8081
	log.Fatal(http.ListenAndServe(":8081", router))
}
// ConnecttoDB : function to connect to mongoDB locally
func ConnecttoDB() *mongo.Collection {

	// Set client options
	//change the URI according to your database
	clientOptions := options.Client().ApplyURI("mongodb+srv://abhitcr1:Abhit1010@cluster0.yg0tb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//DB collection address which we are going to use
	//available to functions of all scope
	collection := client.Database("Appointy").Collection("InstaPosts")

	return collection
}
