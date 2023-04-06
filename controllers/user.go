package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/database"
	"github.com/whicencer/golang-auth/models"
)

// GetUser Get a single user
func GetUser(c *fiber.Ctx) error {
	db := database.DB
	username := c.Params("nickname")
	var user models.User

	result := db.Where("nickname = ?", username).First(&user)

	if result.Error != nil {
		errorMessage := fmt.Sprintf("user %s not found", username)
		return c.JSON(fiber.Map{
			"message":    errorMessage,
			"userExists": false,
		})
	}

	return c.JSON(fiber.Map{
		"message":    "User found",
		"userExists": true,
		"user":       user,
	})
}
