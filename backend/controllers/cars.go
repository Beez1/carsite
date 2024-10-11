package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/beez1/carsite/backend/config"
    "github.com/beez1/carsite/backend/models"
)

// GetCars handles the request to get all cars
func GetCars(w http.ResponseWriter, r *http.Request) {
    var cars []models.Car
    config.DB.Find(&cars)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cars)
}

// GetCar handles the request to get a single car by its ID
func GetCar(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var car models.Car
    if err := config.DB.First(&car, params["id"]).Error; err != nil {
        http.Error(w, "Car not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(car)
}
