package models

// Car represents a car entity
type Car struct {
	Make        string  `json:"make"`
	CarModel    string  `json:"carmodel"`
	Year        int     `json:"year"`
	Price       float64 `json:"price"`
	Mileage     int     `json:"mileage"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
}

