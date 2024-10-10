package model

import (
	"time"
)

type Orders struct {
	OrderId         int       `gorm:"type:int;primary_key"`
	CarID           int       `gorm:"not null"`                  // Car ID as an integer
	OrderDate       time.Time `gorm:"type:date;not null"`        // Date for order creation
	PickupDate      time.Time `gorm:"type:date;not null"`        // Date for pickup
	DropoffDate     time.Time `gorm:"type:date;not null"`        // Date for dropoff
	PickupLocation  string    `gorm:"type:varchar(50);not null"` // varchar(50) equivalent for pickup location
	DropoffLocation string    `gorm:"type:varchar(50);not null"` // varchar(50) equivalent for dropoff location
}

type OrdersRequest struct {
	CarID           string    `json:"car_id" binding:"required"`
	OrderDate       time.Time `json:"order_date" binding:"required"`
	PickupDate      time.Time `json:"pickup_date" binding:"required"`
	DropoffDate     time.Time `json:"dropoff_date" binding:"required"`
	PickupLocation  string    `json:"pickup_location" binding:"required"`
	DropoffLocation string    `json:"dropoff_location" binding:"required"`
}

type OrdersResponse struct {
	CarID           string    `json:"car_id"`
	OrderDate       time.Time `json:"order_date"`
	PickupDate      time.Time `json:"pickup_date"`
	DropoffDate     time.Time `json:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location"`
}
