package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.Background()

func Connect() (*mongo.Database, error) {

	dbUri := os.Getenv("DATABASE_URI")
	dbName := os.Getenv("DATABASE_NAME")

	if dbUri == "" {
		dbUri = "mongodb://localhost:27017"
		log.Println("env DATABASE_URI nggak ada isinya, using local database")
		log.Println(dbUri)
	}

	if dbName == "" {
		dbName = "pdcnotification"
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(dbUri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(Ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("pdcnotification"), nil
}
