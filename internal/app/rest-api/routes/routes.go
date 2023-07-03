package routes

import (
	"net/http"

	"github.com/DmitryLogunov/trading-platform-backend/internal/app/rest-api/controllers"
)

func CreateRoutes(mux *http.ServeMux, c *controllers.Controllers) {
	mux.HandleFunc("/add-alert", c.AddAlert)
}
