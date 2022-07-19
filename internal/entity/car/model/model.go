package model

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	IDCar     int64  `json:"id_car" gorm:"primaryKey"`
	NumberCar string `json:"number_car"`
	Price     int64  `json:"price"`
	Locate    Locate `json:"locate"`
}

type Locate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Area      string  `json:"area"`
	City      string  `json:"city"`
	District  string  `json:"district"`
	Street    string  `json:"street"`
	house     string  `json:"house"`
}
