package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	// custom modules
	"gomodules/CommonUtils"
	"gomodules/databaseUtils"
)

func main() {
	// sample mongo connection string
	// const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"

	envloaderr := godotenv.Load()

	uri := os.Getenv("MONGOURI")

	if envloaderr != nil {
		// panic("Unable to load ENVs" + envloaderr.Error())
		// or
		log.Fatal("Unable to load ENVs" + envloaderr.Error())
	}

	// calling ConnectToMongoDB from databaseUtils
	dbClient, ctx, cancelfunc, err := databaseUtils.ConnectToMongoDB(uri)

	if err != nil {
		panic(err)
	}

	//(defer func will be called just before the parent func(in this case main func) finishes)
	// calling CloseDbConnection to close the db connection
	defer databaseUtils.CloseDbConnection(cancelfunc, dbClient, ctx)

	// Calling module initializer from CommonUtils
	CommonUtils.ModuleInitializer()

	// Calling router initializer from CommonUtils
	CommonUtils.RouterInitializer()

}
