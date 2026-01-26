package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBconnection() *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load variables from .env file, relying on system envs")
	}

	uri := os.Getenv("DB_URL")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("unable to connect to the database:", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("unable to ping the database:", err)
	}

	// Get database name from URI or use default
	dbName := "examAuth"
	// You could extract dbName from URI if needed

	return client.Database(dbName)
}
