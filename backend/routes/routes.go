package routes

import (
	"github.com/gorilla/mux"
	"carsite-backend/controllers"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define routes for cars
	router.HandleFunc("/cars", controllers.GetCars).Methods("GET")

	return router
}
