package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func ConnectDB() (*mongo.Database, error) {
	mongodbHost := os.Getenv("MONGODB_HOST")
	mongodbPort := os.Getenv("MONGODB_PORT")
	mongodbUser := os.Getenv("MONGODB_USER")
	mongodbPassword := os.Getenv("MONGODB_PASSWORD")
	mongodbDatabase := os.Getenv("MONGODB_DATABASE")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongodbUser, mongodbPassword, mongodbHost, mongodbPort)

	fmt.Println(uri)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database(mongodbDatabase)

	fmt.Println("Successfully connected to MongoDB")

	return db, nil
}
