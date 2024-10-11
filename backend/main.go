package main

import (
    "log"
    "net/http"
    "github.com/beez1/carsite/backend/config"
    "github.com/beez1/carsite/backend/routes"
)

func main() {
    // Connect to the database
    config.ConnectDatabase()

    // Initialize routes
    router := routes.InitializeRoutes()

    // Start the server on port 8080
    log.Fatal(http.ListenAndServe(":8080", router))
}
