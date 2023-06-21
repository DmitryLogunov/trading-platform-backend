package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	graphqlApi "github.com/DmitryLogunov/trading-platform/internal/app/graphql-api"
	gqlServices "github.com/DmitryLogunov/trading-platform/internal/app/graphql-api/gql-services"
	"github.com/DmitryLogunov/trading-platform/internal/app/graphql-api/resolvers"
	"github.com/DmitryLogunov/trading-platform/internal/core/database/mongodb"
	"github.com/DmitryLogunov/trading-platform/internal/core/scheduler"
)

const defaultPort = "3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("HTTP_GRAPHQL_PORT")
	if port == "" {
		port = defaultPort
	}

	mongoDB, err := mongodb.ConnectDB()
	if err != nil {
		log.Fatal("Error mongodb connecting")
	}

	scheduler := scheduler.JobsManager{}
	scheduler.Init()

	graphQLServices := gqlServices.GqlServices{}
	graphQLServices.Init()

	var srv = handler.NewDefaultServer(graphqlApi.NewExecutableSchema(graphqlApi.Config{Resolvers: &resolvers.Resolver{
		MongoDB:     mongoDB,
		Scheduler:   &scheduler,
		GqlServices: &graphQLServices,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
