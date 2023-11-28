package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graphqlApi "github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api"
	gqlServices "github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api/gql-services"
	"github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api/resolvers"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/database/mongodb"
	binanceAPIClient "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/scheduler"
	"github.com/go-chi/chi"
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

	router := chi.NewRouter()

	//Add CORS middleware around every request
	//See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	var srv = handler.NewDefaultServer(graphqlApi.NewExecutableSchema(graphqlApi.Config{Resolvers: &resolvers.Resolver{
		MongoDB:          mongoDB,
		Scheduler:        &scheduler,
		GqlServices:      &graphQLServices,
		BinanceAPIClient: &binanceAPIClient.BinanceAPIClient{},
	}}))

	router.Handle("/playground", playground.Handler("GraphQL playground", "/playground"))
	router.Handle("/graphql", srv)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		panic(err)
	}

	log.Printf("connect to http://localhost:%s/graphql for GraphQL queries", port)
	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
}
