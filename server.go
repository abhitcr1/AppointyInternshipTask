package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
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