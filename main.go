package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nelisa-dludla/shopper-api/routes"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")

	router := http.NewServeMux()

	cors := cors.New(cors.Options{
    AllowedOrigins: []string{"https://*", "http://*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    AllowedHeaders: []string{"*"},
    AllowCredentials: false,
    ExposedHeaders: []string{"Link"},
    MaxAge: 300,
	})

	routes.SetupRoutes(router)

	if portString == "" {
		log.Fatal("portString is not available in the environment")
	}

	server := http.Server{
		Addr:    ":" + portString,
		Handler: cors.Handler(router),
	}

	log.Printf("Server starting on port %v\n", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
