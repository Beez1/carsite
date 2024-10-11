package routes

import (
    "github.com/gorilla/mux"
    "/backend/controllersd"
)

func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()

    // Define routes for the cars API
    router.HandleFunc("/cars", controllers.GetCars).Methods("GET")
    router.HandleFunc("/cars/{id}", controllers.GetCar).Methods("GET")

    return router
}
