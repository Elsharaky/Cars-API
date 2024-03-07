package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Elsharaky/Cars-API.git/config"
	"github.com/Elsharaky/Cars-API.git/models"
	"github.com/Elsharaky/Cars-API.git/services"
)

type TypeController interface {
	GetTypes(w http.ResponseWriter, r *http.Request)
	CreateType(w http.ResponseWriter, r *http.Request)
}

type typeController struct {
	service services.TypeService
}

func NewTypeController() *typeController {
	return &typeController{
		service: services.NewTypeService(),
	}
}

func (tc *typeController) CreateType(w http.ResponseWriter, r *http.Request) {
	validator := config.GetValidator()

	w.Header().Set("Content-Type", "application/json")
	var carType models.Type

	if err := json.NewDecoder(r.Body).Decode(&carType); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.Struct(carType); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := tc.service.CreateType(&carType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(carType)
}

func (tc *typeController) GetTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	types, err := tc.service.GetTypes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(types)
}
