package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"carsite-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

var client *mongo.Client

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// Get MongoDB URI and port from environment variables
	mongoURI := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	if mongoURI == "" || port == "" {
		log.Fatalln("Missing required environment variables (DATABASE_URL, PORT)")
	}

	// MongoDB connection setup using the DATABASE_URL from environment variables
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(ctx)

	// Ping MongoDB to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Define HTTP routes
	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			InsertCar(w, r)  // Handle POST request to insert car
		} else if r.Method == http.MethodGet {
			GetCars(w, r)    // Handle GET request to retrieve cars
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// Start the HTTP server using the PORT from environment variables
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// InsertCar handles the request to insert a new car into MongoDB
func InsertCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car

	// Decode the request body into the Car struct
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	collection := client.Database("test").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, car)
	if err != nil {
		http.Error(w, "Failed to insert car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Car inserted successfully!"})
}

// GetCars handles the request to retrieve all cars from MongoDB
func GetCars(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("test").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to retrieve cars", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var cars []models.Car
	if err := cursor.All(ctx, &cars); err != nil {
		http.Error(w, "Error reading cars", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
