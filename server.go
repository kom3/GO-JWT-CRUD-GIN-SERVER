package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type userDetails struct {
	Name   string `json:"name"`
	Age    int8   `json:"age"`
	Skills string `json:"skills"`
}

func getAPI(c *gin.Context) {
	data := []userDetails{
		{Name: "Smith", Age: 25, Skills: "No Skills"},
		{Name: "John", Age: 25, Skills: "Rapper"},
		{Name: "Manju", Age: 25, Skills: "all rounder"},
	}

	c.IndentedJSON(http.StatusOK, data)

}

func postAPI(c *gin.Context) {
	// data := []userDetails{
	// 	{name: "Manju", age: 25, skills: "all rounder"},
	// 	{name: "Smith", age: 25, skills: "No Skills"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// }

	c.IndentedJSON(http.StatusOK, gin.H{"name": "Manju"})
}

func ConnectToMongoDB(uri string) *mongo.Client {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return client
}

func main() {
	// const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"
	const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	// dbClient := ConnectToMongoDB(uri)

	dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = dbClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := dbClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	collectionObj := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	cursor, err := collectionObj.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor...")
		panic(err)
	}

	var rows []bson.D
	fmt.Println("cursor created...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}

	router := gin.Default()
	router.GET("/testget", getAPI)
	router.POST("/testpost", postAPI)
	router.Run("localhost:7777")
}
