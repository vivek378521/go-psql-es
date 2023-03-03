package database

import (
	"fmt"
	"log"

	"example.com/go-psql-es/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "fold"
	dbUser := "fold"
	password := "fold"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Project{})
	DB.AutoMigrate(&models.UserProject{})
	DB.AutoMigrate(&models.Hashtag{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
