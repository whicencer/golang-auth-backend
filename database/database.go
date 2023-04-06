package database

import (
	"log"

	"github.com/whicencer/golang-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dbURL := "postgresql://postgres:MeEUeO04ae4gKTfnnaER@containers-us-west-56.railway.app:6085/railway"

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
