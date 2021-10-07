package database

// Import the libraries
import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Struture to find the data in the database
type AccountVariables struct {
	Email, Password string
}

var Account AccountVariables
var Connect = Connection()

// Connect to Mongodb database
func Connection() *mongo.Client {
	// Use the database URL to connect
	clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@<clustername>.2inap.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// Insert the data in the database
func InsertData(inputData interface{}) {
	// Connect to database and insert the data in it
	collection := Connect.Database("covid").Collection("tracker")
	if _, err := collection.InsertOne(context.TODO(), inputData); err != nil {
		log.Fatal(err)
	}
}

// FindAccount search the data in the database and returns boolean
func FindAccount(myEmail, myPassword string) bool {
	// Connect to database and find the data in it
	collection := Connect.Database("covid").Collection("tracker")
	if err := collection.FindOne(context.TODO(), bson.M{"email": myEmail}).Decode(&Account); err != nil {
		return err == nil
	}
	// Compare the database password with login password
	err := bcrypt.CompareHashAndPassword([]byte(Account.Password), []byte(myPassword))
	return err == nil
}
