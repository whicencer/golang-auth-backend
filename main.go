package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/whicencer/golang-auth/database"
	"github.com/whicencer/golang-auth/routes"
)

func main() {
	// Create a new app
	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// CORS middleware
	app.Use(cors.New())

	// Connecting Database
	database.Connect()

	// Setup routes
	routes.SetupAllRoutes(app)

	// Listen on host
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
