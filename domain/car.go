package domain

import "go-test/model"

type CarRepository interface {
	GetAllCars() ([]model.Cars, error)
	CreateCars(model.Cars) (model.Cars, error)
	UpdateCarsById(model.Cars, int) (model.Cars, error)
	DeleteCarsById(id int) error
}

type CarUseCase interface {
	GetAllCars() ([]model.Cars, error)
	CreateCars(model.CarsRequest) (model.Cars, error)
	UpdateCarsById(model.CarsRequest, int) (model.Cars, error)
	DeleteCarsById(id int) error
}
