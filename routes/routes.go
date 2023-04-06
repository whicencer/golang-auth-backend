package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/routes/auth"
	"github.com/whicencer/golang-auth/routes/user"
)

// To fix
func SetupAllRoutes(app *fiber.App) {
	user.SetupRoutes(app)
	auth.SetupRoutes(app)
}
