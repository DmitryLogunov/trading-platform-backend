package main

import (
	"github.com/DmitryLogunov/trading-platform/internal/api/graphql-api"
	"github.com/DmitryLogunov/trading-platform/internal/api/graphql-api/resolvers"
	"github.com/DmitryLogunov/trading-platform/internal/database/mysql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultPort
	}

	// establish connection
	mysql.ConnectDB()
	// create db
	mysql.CreateDB()
	// migrate the db with Post models
	mysql.MigrateDB()

	var srv = handler.NewDefaultServer(graphql_api.NewExecutableSchema(graphql_api.Config{Resolvers: &resolvers.Resolver{
		Database: mysql.DBInstance,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
