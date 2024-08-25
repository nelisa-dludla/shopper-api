package routes

import (
	"net/http"

	"github.com/nelisa-dludla/shopper-api/handlers"
)

func SetupRoutes(router *http.ServeMux) {
  router.HandleFunc("GET /products", handlers.GetProducts)
}
