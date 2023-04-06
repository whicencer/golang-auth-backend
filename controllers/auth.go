package controllers

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/database"
	"github.com/whicencer/golang-auth/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	// init a "db" and "user" variable
	db := database.DB

	var body struct {
		Nickname    string
		Description string
		Password    string
	}

	// Getting an error on parsing a request body data
	// If there is an error - return a StatusBadRequest with specified JSON
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"ok":      false,
		})
	}

	// Validate body parameters
	if len(body.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password length should be 8 symbols or more",
			"ok":      false,
		})
	}

	if len(body.Nickname) < 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username length should be 2 symbols or more",
			"ok":      false,
		})
	}

	// Define a hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	// If there are some errors by hashing a password - return a Status Internal Error with JSON
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
			"ok":      false,
		})
	}

	user := models.User{Nickname: body.Nickname, Description: body.Description, Password: string(hashedPassword)}

	if err := db.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicated key not allowed") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Nickname already taken",
				"ok":      false,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created",
		"user":    user,
		"ok":      true,
	})
}

func Login(c *fiber.Ctx) error {
	db := database.DB

	var body struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"ok":      false,
		})
	}

	var dbUser models.User
	if err := db.Where("nickname = ?", body.Nickname).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid login or password!",
				"ok":      false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user",
			"ok":      false,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid login or password!",
			"ok":      false,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["nickname"] = dbUser.Nickname
	claims["description"] = dbUser.Description
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "You have successfully logged in",
		"token":   t,
		"ok":      true,
	})
}

// GetMe method
func GetMe(c *fiber.Ctx) error {
	// Getting authorization header and check if it's not empty
	authHeader := c.Get("Authorization")
	if len(authHeader) <= len("Bearer ") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Authorization token is missing or invalid",
			"ok":      false,
		})
	}

	tokenString := authHeader[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Устанавливаем ключ для проверки подлинности JWT-токена
		return []byte(SecretKey), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
			"ok":      false,
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		nickname := claims["nickname"].(string)
		description := claims["description"].(string)

		return c.JSON(fiber.Map{
			"nickname":    nickname,
			"description": description,
			"ok":          true,
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong token",
			"ok":      false,
		})
	}

}
