package usecase

import (
	"go-test/domain"
	"go-test/model"
)

func NewCarUseCase(carRepository domain.CarRepository) domain.CarUseCase {
	return &carUseCase{carRepository: carRepository}
}

type carUseCase struct {
	carRepository domain.CarRepository
}

func (c *carUseCase) CreateCars(carsRequest model.CarsRequest) (model.Cars, error) {
	cars := model.Cars{
		CarName:   carsRequest.CarName,
		DayRate:   carsRequest.DayRate,
		MonthRate: carsRequest.MonthRate,
		Image:     carsRequest.Image,
	}

	cars, err := c.carRepository.CreateCars(cars)
	return cars, err
}

// DeleteCarsById implements domain.CarUseCase.
func (c *carUseCase) DeleteCarsById(id int) error {
	err := c.carRepository.DeleteCarsById(id)
	return err
}

// UpdateCars implements domain.CarUseCase.
func (c *carUseCase) UpdateCarsById(carsRequest model.CarsRequest, id int) (model.Cars, error) {
	cars := model.Cars{
		CarName:   carsRequest.CarName,
		DayRate:   carsRequest.DayRate,
		MonthRate: carsRequest.MonthRate,
		Image:     carsRequest.Image,
	}

	carsResp, err := c.carRepository.UpdateCarsById(cars, id)
	return carsResp, err
}

// GetAllCars implements domain.CarUseCase.
func (c *carUseCase) GetAllCars() ([]model.Cars, error) {
	cars, err := c.carRepository.GetAllCars()
	return cars, err
}
