package model

type Cars struct {
	CarId     int     `gorm:"type:int;primary_key"`
	CarName   string  `gorm:"type:varchar(50);not null"`  // varchar(50) equivalent
	DayRate   float64 `gorm:"not null"`                   // double equivalent
	MonthRate float64 `gorm:"not null"`                   // double equivalent
	Image     string  `gorm:"type:varchar(256);not null"` // varchar(256) equivalent
}

type CarsRequest struct {
	CarName   string  `json:"car_name" binding:"required"`
	DayRate   float64 `json:"day_rate" binding:"required"`
	MonthRate float64 `json:"month_rate" binding:"required"`
	Image     string  `json:"image" binding:"required"`
}

type CarsResponse struct {
	CarId     int     `json:"car_id"`
	CarName   string  `json:"car_name"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
}
