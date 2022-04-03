package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// custom modules
	"gomodules/APIS"
)

// global variables

func postAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"name": "Manju"})
}

func main() {
	// const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"
	const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	// dbClient, ctx, cancelfunc, err := ConnectToMongoDB(uri)
	dbClient, ctx, cancelfunc, err := APIS.ConnectToMongoDB(uri)

	if err != nil {
		panic(err)
	}

	// function to close the db connection
	defer APIS.CloseDbConnection(cancelfunc, dbClient, ctx)

	// API routes and handlers
	router := gin.Default()
	router.GET("/getemployeelist", APIS.GetEmployeeList)
	router.POST("/testpost", postAPI)
	router.Run("localhost:7777")
}
