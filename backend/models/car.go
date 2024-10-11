package models

import "gorm.io/gorm"

type Car struct {
    gorm.Model
    Make        string  `json:"make"`
    Model       string  `json:"model"`
    Year        int     `json:"year"`
    Price       float64 `json:"price"`
    Mileage     int     `json:"mileage"`
    Description string  `json:"description"`
    ImageURL    string  `json:"image_url"`
}

