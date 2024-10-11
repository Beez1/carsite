package controllers

import (
	"encoding/json"
	"net/http"

	"carsite-backend/models"
	"carsite-backend/config"
)

// GetCars handles the request to retrieve all cars
func GetCars(w http.ResponseWriter, r *http.Request) {
	var cars []models.Car
	config.DB.Find(&cars)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

