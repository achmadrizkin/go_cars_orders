package repo

import (
	"errors"
	"go-test/domain"
	"go-test/model"

	"gorm.io/gorm"
)

func NewCarRepository(db *gorm.DB) domain.CarRepository {
	return &carRepository{db: db}
}

type carRepository struct {
	db *gorm.DB
}

func (c *carRepository) CreateCars(cars model.Cars) (model.Cars, error) {
	err := c.db.Create(&cars).Error
	if err != nil {
		return cars, errors.New("internal server error: " + err.Error())
	}
	return cars, nil
}

func (c *carRepository) DeleteCarsById(id int) error {
	if err := c.db.Delete(&model.Cars{}, id).Error; err != nil {
		return errors.New("failed to delete car: " + err.Error())
	}
	return nil
}

func (c *carRepository) UpdateCarsById(car model.Cars, id int) (model.Cars, error) {
	// Find the car by ID and update with new data
	err := c.db.Model(&model.Cars{}).Where("car_id = ?", id).Updates(car).Error
	if err != nil {
		return car, errors.New("failed to update car: " + err.Error())
	}
	return car, nil
}

func (g *carRepository) GetAllCars() ([]model.Cars, error) {
	var cars []model.Cars

	err := g.db.Model(model.Cars{}).Find(&cars).Error
	if err != nil {
		return cars, errors.New("internal server error: " + err.Error())
	}
	return cars, nil
}
