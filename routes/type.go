package routes

import (
	"github.com/Elsharaky/Cars-API.git/controllers"
	"github.com/gorilla/mux"
)

func SetupTypeRoutes(router *mux.Router) {
	typeController := controllers.NewTypeController()
	router.HandleFunc("/api/v1/types", typeController.GetTypes).Methods("GET")
	router.HandleFunc("/api/v1/types", typeController.CreateType).Methods("POST")
}
