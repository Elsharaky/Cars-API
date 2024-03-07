package routes

import (
	"github.com/Elsharaky/Cars-API.git/controllers"
	"github.com/gorilla/mux"
)

func SetupCarsRoutes(router *mux.Router) {
	carController := controllers.NewCarController()
	router.HandleFunc("/api/v1/cars", carController.GetCars).Methods("GET")
	router.HandleFunc("/api/v1/cars", carController.CreateCar).Methods("POST")
	router.HandleFunc("/api/v1/cars/{id}", carController.GetCarByID).Methods("GET")
}
