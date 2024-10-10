package usecase

import (
	"go-test/domain"
	"go-test/model"
)

func NewOrderUseCase(orderRepository domain.OrderRepository) domain.OrderUseCase {
	return &OrderUseCase{orderRepository: orderRepository}
}

type OrderUseCase struct {
	orderRepository domain.OrderRepository
}

// CreateOrders implements domain.OrderUseCase.
func (*OrderUseCase) CreateOrders(model.OrdersRequest) (model.Orders, error) {
	panic("unimplemented")
}

// DeleteOrdersById implements domain.OrderUseCase.
func (*OrderUseCase) DeleteOrdersById(id int) error {
	panic("unimplemented")
}

// GetAllOrders implements domain.OrderUseCase.
func (o *OrderUseCase) GetAllOrders() ([]model.Orders, error) {
	orderList, err := o.orderRepository.GetAllOrders()
	return orderList, err
}

// UpdateOrdersById implements domain.OrderUseCase.
func (*OrderUseCase) UpdateOrdersById(model.OrdersRequest, int) (model.Orders, error) {
	panic("unimplemented")
}
