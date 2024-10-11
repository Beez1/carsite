package config

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection (no return value)
func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using system environment variables instead")
	}

	// Get the database connection string from environment variable
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is required but not set")
	}

	// Open a connection to the PostgreSQL database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Assign database connection to the global DB variable
	DB = database

	// Log successful connection
	log.Println("Connected to the database successfully!")
}
