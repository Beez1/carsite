package main

import (
    "log"
    "carsite-backend/cloudsql"
    "github.com/joho/godotenv"
)

func main() {
    // Load the environment variables from the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalln("Error loading .env file")
    }

    // Connect to the Cloud SQL instance
    db, err := cloudsql.ConnectWithConnector()
    if err != nil {
        log.Fatalln("Failed to connect to the database:", err)
    }
    defer db.Close()

    log.Println("Application started successfully!")
}
