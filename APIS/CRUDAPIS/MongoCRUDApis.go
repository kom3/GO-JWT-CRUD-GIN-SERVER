package CRUDAPIS

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/bson/primitive"

	// Validation package
	"github.com/go-playground/validator/v10"

	// custom modules
	"gomodules/databaseUtils"
	"gomodules/models"
	"gomodules/responses"
)

// global declarations
var mongo_db_con_obj databaseUtils.MongoDbConObjectType
var dbClient *mongo.Client
var validate = validator.New()

type EmployeeDetails models.EmployeeDetails

// Methods
func Init() {
	fmt.Println("<== crudApis.go ===> ")
	mongo_db_con_obj = databaseUtils.MongoDbConObj
	dbClient = mongo_db_con_obj.Client
}

func SamplePostAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "This is Sample Post API"})
}

func SampleGETAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "This is Sample GET API"})
}

func GetAllEmployeesList(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []EmployeeDetails{
	// 	{Name: "Manju", Age: 25, Skills: "all rounder"},
	// 	{Name: "Smith", Age: 25, Skills: "No Skills"},
	// 	{Name: "John", Age: 25, Skills: "Rapper"},
	// 	{Name: "John", Age: 25, Skills: "Rapper"},
	// 	{Name: "John", Age: 25, Skills: "Rapper"},
	// }

	// Select db and collection
	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	filter := bson.D{}

	// Fetch data from collection
	cursor, err := employeeCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error while creating cursor in GetEmployeeList...", err)
		panic(err)
	}

	// initialising an empty array of type bson.D
	// var rows []bson.D

	// initialising an empty array of type EmployeeDetails
	var rows []EmployeeDetails
	// or
	// var rows = []EmployeeDetails{}

	fmt.Println("Cursor created in GetEmployeeList...")
	if err = cursor.All(context.TODO(), &rows); err != nil {
		fmt.Println("Error while fetching rows in GetEmployeeList...")
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}

	// Simple response
	// c.JSON(200, gin.H{
	// 	"data": "Hello from Gin-gonic & mongoDB",
	// })

	// initialising response type as an
	var result responses.GetAllEmployeesResp
	result.Error.Code = 0
	result.Error.Message = "Success"
	result.Results = rows

	c.IndentedJSON(http.StatusOK, result)

}

func GetEmployeeById(c *gin.Context) {
	fmt.Println("GetEmployeeById...")

	//Below method is to read url param,
	//Example 1: ip:port/myurl/:id ==> ip:port/myurl/1234,
	//Example 2: ip:port/myurl/:empname/:age/:mobile ==> ip:port/myurl/manju/24/999999999999,
	var empname = c.Param("empname")
	fmt.Println("URL param --> emp_id: ", empname)

	//Below methods is to read query param,
	//Example: ip:port/myurl ==> ip:port/myurl?age=24,
	var age = c.Query("age")
	fmt.Println("Query param --> age: ", age)

	//c.GetQuery returns Query param value and bool(true, if exits else false)
	var mobile, exists = c.GetQuery("mobile") // shortcut for c.Request.URL.Query().Get(key)
	fmt.Println("Query param --> skill: ", mobile, exists)

	//Below method is to read data from request body
	type sampleRequestBodyDataStructure struct {
		Name   string
		Age    int8
		Mobile int64
	}
	var varToStoreRequestData sampleRequestBodyDataStructure
	// Binding JSON data from request body to variable
	if bindJsonErr := c.BindJSON(&varToStoreRequestData); bindJsonErr != nil {
		fmt.Println("Bind JSON Error")
	}
	fmt.Println("This is the data from request body(BindJSON):", varToStoreRequestData)

	//Using c.GetRawData() to read JSON data from request body
	jsonData, _ := c.GetRawData()
	fmt.Println("This is the data from request body(GetRawData):", jsonData)

	// 	-------------------------------------|
	//	Database Operation Starts from here	 |
	//	-------------------------------------|

	// Fetch Employee from database based on employee name
	emp_name := c.Query("name")
	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	// D is an ordered representation of a BSON document.
	// This type should be used when the order of the elements matters,
	// such as MongoDB command documents.
	// If the order of the elements does not matter, an M should be used instead.

	// A D should not be constructed with duplicate key names, as that can cause undefined server behavior.

	// Example usage:
	// bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}} //{"key", "value"} here key must be unique
	// bson.M{"foo": "bar"}

	filter := bson.M{"name": emp_name} //bson-->binary javascript object notation

	var result responses.EmployeeDetailsNode
	var response responses.GetEmployeeByIdResp

	err := employeeCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error while fetching data in GetEmployeeById...for id=", emp_name, err)
		// panic(err)
		response.Error.Code = 404
		response.Error.Message = "No records found"
		response.Results = result
		// response.Results = []interface{}{} // this way can assign an empty array as value(for this Results node must be type of interface{} in struct GetEmployeeByIdResp )
	} else {
		response.Error.Code = 0
		response.Error.Message = "Success"
		response.Results = result
	}

	c.IndentedJSON(http.StatusOK, response)

}

func AddEmployee(c *gin.Context) {
	fmt.Println("AddEmployee...")

	employeeCollection := dbClient.Database("testdatabase").Collection("employee")

	var employee models.EmployeeDetails
	var response responses.AddEmployee

	//validate the json data from request body against employee details model
	if err := c.BindJSON(&employee); err != nil {
		response.Error.Code = 1
		response.Error.Message = "Failed to add employee, Invalid JSON data received!"
		c.IndentedJSON(http.StatusOK, response)
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&employee); validationErr != nil {
		response.Error.Code = 1
		response.Error.Message = "Failed to add employee, Validation failed!"
		c.IndentedJSON(http.StatusOK, response)
		return
	}

	//  Adding employee in to collection
	result, err := employeeCollection.InsertOne(context.TODO(), employee)
	if err != nil {
		response.Error.Code = 1
		response.Error.Message = "Failed to add employee, DB insert error!"
		c.IndentedJSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, result)

}

func AddEmployees(c *gin.Context) {
	fmt.Println("GetEmployeeList...")

	// harcoded data
	// rows := []EmployeeDetails{
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
	var rows []EmployeeDetails
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
	// rows := []EmployeeDetails{
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
	var rows []EmployeeDetails
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
	// rows := []EmployeeDetails{
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
	var rows []EmployeeDetails
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
	// rows := []EmployeeDetails{
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
	var rows []EmployeeDetails
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
	// rows := []EmployeeDetails{
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
	var rows []EmployeeDetails
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
