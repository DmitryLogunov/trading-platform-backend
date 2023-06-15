package main

import (
	graphql_api "github.com/DmitryLogunov/trading-platform/core/graphql-api"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DmitryLogunov/trading-platform/core/database"
	"github.com/DmitryLogunov/trading-platform/core/graphql-api/resolvers"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultPort
	}

	// establish connection
	database.ConnectDB()
	// create db
	database.CreateDB()
	// migrate the db with Post models
	database.MigrateDB()

	var srv = handler.NewDefaultServer(graphql_api.NewExecutableSchema(graphql_api.Config{Resolvers: &resolvers.Resolver{
		Database: database.DBInstance,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
