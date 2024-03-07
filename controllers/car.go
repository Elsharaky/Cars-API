package controllers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/Elsharaky/Cars-API.git/config"
	"github.com/Elsharaky/Cars-API.git/models"
	"github.com/Elsharaky/Cars-API.git/services"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

type CarController interface {
	CreateCar(w http.ResponseWriter, r *http.Request)
	GetCars(w http.ResponseWriter, r *http.Request)
	GetCarByID(w http.ResponseWriter, r *http.Request)
}

type carController struct {
	service services.CarService
}

func NewCarController() *carController {
	return &carController{
		service: services.NewCarService(),
	}
}

func (cc *carController) CreateCar(w http.ResponseWriter, r *http.Request) {
	validator := config.GetValidator()

	w.Header().Set("Content-Type", "application/json")
	var car models.Car

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.Struct(car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := cc.service.CreateCar(&car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func (cc *carController) GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var car models.Car
	queries := r.URL.Query()

	if queries.Get("make") != "" {
		car.Make = queries.Get("make")
	}

	if queries.Get("color") != "" {
		car.Color = queries.Get("color")
	}

	if queries.Get("type") != "" {
		car.Type = queries.Get("type")
	}

	if queries.Get("speed") != "" {
		var spped pq.Int32Array
		if err := json.Unmarshal([]byte(queries.Get("speed")), &spped); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		car.Speed = spped
	}

	if queries.Get("model") != "" {
		model, err := strconv.ParseUint(queries.Get("model"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		car.Model = uint(model)
	}

	if !reflect.DeepEqual(car, reflect.Zero(reflect.TypeOf(car))) {
		cars, err := cc.service.GetCarsByFilter(car)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(cars)
		return
	}

	cars, err := cc.service.GetCars()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cars)
}

func (cc *carController) GetCarByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car, err := cc.service.GetCarByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(car)
}
