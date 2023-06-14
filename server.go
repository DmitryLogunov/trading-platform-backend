package main

import (
	"github.com/DmitryLogunov/golang-graphql/database"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DmitryLogunov/golang-graphql/graph"
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
	// migrate the db with Post model
	database.MigrateDB()

	var srv = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Database: database.DBInstance,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
