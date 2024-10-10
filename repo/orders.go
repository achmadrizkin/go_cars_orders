package repo

import (
	"errors"
	"go-test/domain"
	"go-test/model"

	"gorm.io/gorm"
)

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{db: db}
}

type orderRepository struct {
	db *gorm.DB
}

// CreateOrders implements domain.OrderRepository.
func (*orderRepository) CreateOrders(model.Orders) (model.Orders, error) {
	panic("unimplemented")
}

// DeleteOrdersById implements domain.OrderRepository.
func (*orderRepository) DeleteOrdersById(id int) error {
	panic("unimplemented")
}

// GetAllOrders implements domain.OrderRepository.
func (o *orderRepository) GetAllOrders() ([]model.Orders, error) {
	var orders []model.Orders

	err := o.db.Model(model.Orders{}).Find(&orders).Error
	if err != nil {
		return orders, errors.New("internal server error: " + err.Error())
	}
	return orders, nil
}

// UpdateOrdersById implements domain.OrderRepository.
func (*orderRepository) UpdateOrdersById(model.Orders, int) (model.Orders, error) {
	panic("unimplemented")
}
