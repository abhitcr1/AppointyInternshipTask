package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
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
	router.GET("/User", getUserPosts)
	router.GET("/User/:id", getUser)
	router.GET("/user/search?q=title", searchUser)
	router.POST("/user", createPost)
	router.POST("/user", createUser)
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
//Function to get all user in DataBase
func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	// User array
	var user []User

	// bson.M{},  we passed empty filter of unordered map.
	cur, err := collection.Find(context.TODO(), bson.M{})

	//Error Handling
	if err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	//Loops over the cursor stream and appends to []User array
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var User User
		// decode similar to deserialize process.
		err := cur.Decode(&User)

		//Error Handling
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		user = append(user, User)
	}
	//Error Handling
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Encoding the data in Array to JSON format
	json.NewEncoder(w).Encode(user)

}
//Function to create a new user in Database
func createuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var user user

	//Decoding the Data from JSON format to user variable
	_ = json.NewDecoder(r.Body).Decode(&user)
	//inserts the data from decoded var to MongoDB in BSON format
	result, err := collection.InsertOne(context.TODO(), user)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}
//Function to search User by ID
func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var User User

	// string to primitive.ObjectID (typeCasting)
	id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))

	// creating filter of unordered map with ID as input
	filter := bson.M{"_id": id}

	//Searching in DB with given ID as keyword
	err := collection.FindOne(context.TODO(), filter).Decode(&User)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(User)
}

fmt.Fprintf(w, `Hello world`)
func searchUserUsingID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var article Article

	//recovers the argument of search query present in URL after "q"
	title := string(r.URL.Query().Get("q"))

	//makes an unordered map filter of title
	filter := bson.M{"title": title}

	//Searching in DB with given title as keyword
	err := collection.FindOne(context.TODO(), filter).Decode(&article)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(article)
}