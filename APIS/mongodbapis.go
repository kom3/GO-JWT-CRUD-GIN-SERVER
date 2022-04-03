package APIS

import (
	"fmt"

	"net/http"

	"time"

	"github.com/gin-gonic/gin"

	"context"

	// mongodb packages
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// global variables
var dbClient *mongo.Client

type employeeDetails struct {
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Skills  string `json:"skills"`
	Address string `json:"address"`
}

func ConnectToMongoDB(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	dbClient = client
	fmt.Println("Successfully connected and pinged.")

	return client, ctx, cancel, err
}

func CloseDbConnection(cancel context.CancelFunc, dbClient *mongo.Client, ctx context.Context) {
	// client provides a method to close
	// a mongoDB connection.
	defer cancel()
	defer func() {
		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func GetEmployeeList(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []employeeDetails{
	// 	{name: "Manju", age: 25, skills: "all rounder"},
	// 	{name: "Smith", age: 25, skills: "No Skills"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// }

	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	cursor, err := employeeCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor in GetEmployeeList...", err)
		panic(err)
	}

	// initialising an empty array of type bson.D
	// var rows []bson.D

	// initialising an empty array of type employeeDetails
	var rows []employeeDetails
	// or
	// var rows = []employeeDetails{}

	fmt.Println("Cursor created in GetEmployeeList...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows in GetEmployeeList...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	c.IndentedJSON(http.StatusOK, rows)

}

func GetEmployeeById(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []employeeDetails{
	// 	{name: "Manju", age: 25, skills: "all rounder"},
	// 	{name: "Smith", age: 25, skills: "No Skills"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// }

	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	cursor, err := employeeCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor in GetEmployeeList...", err)
		panic(err)
	}

	// var rows []bson.D
	var rows []employeeDetails
	fmt.Println("Cursor created in GetEmployeeList...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows in GetEmployeeList...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	c.IndentedJSON(http.StatusOK, rows)

}

func AddEmployee(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []employeeDetails{
	// 	{name: "Manju", age: 25, skills: "all rounder"},
	// 	{name: "Smith", age: 25, skills: "No Skills"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// }

	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	cursor, err := employeeCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor in GetEmployeeList...", err)
		panic(err)
	}

	// var rows []bson.D
	var rows []employeeDetails
	fmt.Println("Cursor created in GetEmployeeList...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows in GetEmployeeList...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	c.IndentedJSON(http.StatusOK, rows)

}

func UpdateAllEmployees(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []employeeDetails{
	// 	{name: "Manju", age: 25, skills: "all rounder"},
	// 	{name: "Smith", age: 25, skills: "No Skills"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// }

	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	cursor, err := employeeCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor in GetEmployeeList...", err)
		panic(err)
	}

	// var rows []bson.D
	var rows []employeeDetails
	fmt.Println("Cursor created in GetEmployeeList...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows in GetEmployeeList...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	c.IndentedJSON(http.StatusOK, rows)

}

func UpdateEmployeeById(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []employeeDetails{
	// 	{name: "Manju", age: 25, skills: "all rounder"},
	// 	{name: "Smith", age: 25, skills: "No Skills"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// }

	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	cursor, err := employeeCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor in GetEmployeeList...", err)
		panic(err)
	}

	// var rows []bson.D
	var rows []employeeDetails
	fmt.Println("Cursor created in GetEmployeeList...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows in GetEmployeeList...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	c.IndentedJSON(http.StatusOK, rows)

}

func DeleteEmployeeById(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []employeeDetails{
	// 	{name: "Manju", age: 25, skills: "all rounder"},
	// 	{name: "Smith", age: 25, skills: "No Skills"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// }

	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	cursor, err := employeeCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor in GetEmployeeList...", err)
		panic(err)
	}

	// var rows []bson.D
	var rows []employeeDetails
	fmt.Println("Cursor created in GetEmployeeList...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows in GetEmployeeList...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	c.IndentedJSON(http.StatusOK, rows)

}

func DeleteAllEmployees(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []employeeDetails{
	// 	{name: "Manju", age: 25, skills: "all rounder"},
	// 	{name: "Smith", age: 25, skills: "No Skills"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// 	{name: "John", age: 25, skills: "Rapper"},
	// }

	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	cursor, err := employeeCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor in GetEmployeeList...", err)
		panic(err)
	}

	// var rows []bson.D
	var rows []employeeDetails
	fmt.Println("Cursor created in GetEmployeeList...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows in GetEmployeeList...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	c.IndentedJSON(http.StatusOK, rows)

}
