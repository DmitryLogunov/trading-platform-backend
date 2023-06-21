package rest_api

import (
	"fmt"
	"github.com/DmitryLogunov/trading-platform/internal/app/rest-api/controllers"
	"github.com/DmitryLogunov/trading-platform/internal/app/rest-api/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
)

const defaultPort = "4000"

type Server struct {
	Controllers *controllers.Controllers
}

func NewServer(db *mongo.Database) *Server {
	return &Server{
		Controllers: &controllers.Controllers{
			MongoDB: db,
		},
	}
}

func (s *Server) Start() {
	port := os.Getenv("HTTP_REST_PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()
	routes.CreateRoutes(mux, s.Controllers)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatal(err)
	}

	log.Printf("REST API listening to http://localhost:%s/...", port)
}
