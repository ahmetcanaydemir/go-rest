package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ahmetcanaydemir/go-rest/pkg/api/controller"
	"github.com/ahmetcanaydemir/go-rest/pkg/configs"
)

func main() {
	// Configuration
	connStr := os.Getenv("MONGO_URI")
	port := os.Getenv("PORT")

	if connStr == "" {
		panic(`required environment variable MONGO_URI not set.\nPlease start like following MONGO_URI=uri go run main.go`)
	}
	if port == "" {
		port = "8080"
	}

	configs.Server.Config.DbConnectionString = connStr
	configs.Server.Config.Port = port

	// Controllers
	controller.NewInMemoryController()
	controller.NewMongoController()

	// Server
	fmt.Println("⚡ HTTP server is running at http://localhost:" + configs.Server.Config.Port)
	log.Fatal(http.ListenAndServe(":"+configs.Server.Config.Port, nil))
}
