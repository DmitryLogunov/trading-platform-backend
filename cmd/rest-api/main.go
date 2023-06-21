package main

import (
	restApi "github.com/DmitryLogunov/trading-platform/internal/app/rest-api"
	"github.com/DmitryLogunov/trading-platform/internal/core/database/mongodb"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoDB, err := mongodb.ConnectDB()
	if err != nil {
		log.Fatal("Error mongodb connecting")
	}

	server := restApi.NewServer(mongoDB)
	server.Start()
}
