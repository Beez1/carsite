package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get the database connection string from the environment variable
    dsn := os.Getenv("DB_DSN")
    
    // Open a connection to the PostgreSQL database
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database: ", err)
    }

    // Automatically migrate the schema
    database.AutoMigrate(&models.Car{})

    DB = database
}
