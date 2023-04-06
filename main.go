package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/whicencer/golang-auth/database"
	"github.com/whicencer/golang-auth/routes"
)

func main() {
	// Create a new app
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New())

	// Connecting Database
	database.Connect()

	// Setup routes
	routes.SetupAllRoutes(app)

	// Listen localhost:2000
	app.Listen("localhost:2000")
}
