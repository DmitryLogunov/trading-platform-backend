package main

import (
	"github.com/DmitryLogunov/trading-platform/internal/api/graphql-api"
	"github.com/DmitryLogunov/trading-platform/internal/api/graphql-api/resolvers"
	"github.com/DmitryLogunov/trading-platform/internal/core/scheduler"
	"github.com/DmitryLogunov/trading-platform/internal/database/mongodb"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultPort
	}

	mongoDB, err := mongodb.ConnectDB()
	if err != nil {
		log.Fatal("Error mongodb connecting")
	}

	scheduler := scheduler.JobsManager{}
	scheduler.Init()

	var srv = handler.NewDefaultServer(graphql_api.NewExecutableSchema(graphql_api.Config{Resolvers: &resolvers.Resolver{
		MongoDB:   mongoDB,
		Scheduler: &scheduler,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
