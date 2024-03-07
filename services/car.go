package services

import (
	"github.com/Elsharaky/Cars-API.git/models"
	"github.com/Elsharaky/Cars-API.git/repositories"
)

type CarService interface {
	CreateCar(car *models.Car) error
	GetCars() ([]*models.Car, error)
	GetCarByID(id uint) (*models.Car, error)
	GetCarsByFilter(filter models.Car) ([]*models.Car, error)
}
type carService struct {
	repo repositories.CarRepository
}

func NewCarService() *carService {
	return &carService{
		repo: repositories.NewCarRepository(),
	}
}

func (cs *carService) CreateCar(car *models.Car) error {
	return cs.repo.CreateCar(car)
}

func (cs *carService) GetCars() ([]*models.Car, error) {
	return cs.repo.GetCars()
}

func (cs *carService) GetCarByID(id uint) (*models.Car, error) {
	return cs.repo.GetCarByID(id)
}

func (cs *carService) GetCarsByFilter(filter models.Car) ([]*models.Car, error) {
	return cs.repo.GetCarsByFilter(filter)
}
