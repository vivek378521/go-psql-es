package database

import (
	"fmt"
	"log"
	"os"

	"example.com/go-psql-es/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	dbUser := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PWD")
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
	DB.AutoMigrate(&models.HashtagProject{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
