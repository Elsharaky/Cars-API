package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Elsharaky/Cars-API.git/config"
	"github.com/Elsharaky/Cars-API.git/models"
)

type CarRepository interface {
	CreateCar(car *models.Car) error
	GetCars() ([]*models.Car, error)
	GetCarByID(id uint) (*models.Car, error)
	GetCarsByFilter(filter models.Car) ([]*models.Car, error)
}

type carRepository struct {
	db *sql.DB
}

func NewCarRepository() *carRepository {
	return &carRepository{
		db: config.GetDB(),
	}
}

func (cr *carRepository) CreateCar(car *models.Car) error {
	query := `INSERT INTO cars (name, make, model, color, speed, type_id) 
	VALUES($1,$2,$3,$4,$5,$6)
	 RETURNING id, name, make, model, color, speed, type_id`

	return cr.db.QueryRow(query, car.Name, car.Make, car.Model, car.Color, car.Speed, car.TypeID).Scan(&car.ID, &car.Name, &car.Make, &car.Model, &car.Color, &car.Speed, &car.TypeID)
}

func (cr *carRepository) GetCars() ([]*models.Car, error) {
	query := `SELECT cars.id, cars.name, make, model, color, speed, types.name as type FROM cars JOIN types ON cars.type_id = types.id`

	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []*models.Car{}
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Name, &car.Make, &car.Model, &car.Color, &car.Speed, &car.Type); err != nil {
			return nil, err
		}
		cars = append(cars, &car)
	}

	return cars, nil
}

func (cr *carRepository) GetCarByID(id uint) (*models.Car, error) {
	query := `SELECT cars.id, cars.name, make, model, color, speed, types.name as type FROM cars JOIN types ON cars.type_id = types.id WHERE cars.id = $1`

	var car models.Car
	err := cr.db.QueryRow(query, id).Scan(&car.ID, &car.Name, &car.Make, &car.Model, &car.Color, &car.Speed, &car.Type)
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (cr *carRepository) GetCarsByFilter(filter models.Car) ([]*models.Car, error) {
	query := `SELECT cars.id, cars.name, make, model, color, speed, types.name as type FROM cars JOIN types ON cars.type_id = types.id WHERE cars.id IS NOT NULL`
	filters := make([]interface{}, 0)
	idx := 1

	if filter.Make != "" {
		query += fmt.Sprintf(` AND make = $%d`, idx)
		idx++
		filters = append(filters, filter.Make)
	}
	if filter.Model != 0 {
		query += fmt.Sprintf(` AND model = $%d`, idx)
		idx++
		filters = append(filters, filter.Model)
	}
	if filter.Color != "" {
		query += fmt.Sprintf(` AND color = $%d`, idx)
		idx++
		filters = append(filters, filter.Color)
	}
	if filter.Type != "" {
		query += fmt.Sprintf(` AND types.name = $%d`, idx)
		idx++
		filters = append(filters, filter.Type)
	}
	if filter.Speed != nil {
		query += fmt.Sprintf(` AND speed = $%d`, idx)
		idx++
		filters = append(filters, filter.Speed)
	}

	rows, err := cr.db.Query(query, filters...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cars := []*models.Car{}
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Name, &car.Make, &car.Model, &car.Color, &car.Speed, &car.Type); err != nil {
			return nil, err
		}
		cars = append(cars, &car)
	}
	return cars, nil
}
