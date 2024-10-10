package domain

import "go-test/model"

type OrderRepository interface {
	GetAllOrders() ([]model.Orders, error)
	CreateOrders(model.Orders) (model.Orders, error)
	UpdateOrdersById(model.Orders, int) (model.Orders, error)
	DeleteOrdersById(id int) error
}

type OrderUseCase interface {
	GetAllOrders() ([]model.Orders, error)
	CreateOrders(model.OrdersRequest) (model.Orders, error)
	UpdateOrdersById(model.OrdersRequest, int) (model.Orders, error)
	DeleteOrdersById(id int) error
}
