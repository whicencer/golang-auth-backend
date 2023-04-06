package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)
	app.Get("/auth/me", controllers.GetMe)
}
