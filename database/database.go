package database

import (
	"log"
	"os"

	"github.com/whicencer/golang-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Failed to connect to database, \n", err)
	}
	log.Println("Connected")

	db.AutoMigrate(&models.User{})

	DB = db
}
