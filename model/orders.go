package model

import (
	"github.com/jinzhu/gorm"
)

type Coordinate struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Order struct {
	gorm.Model
	OrderId string `json:"orderId"`
	Username string `json:"username"`
	PickUpLoc []*Coordinate `json:"pick_sup_loc"`
	DropOffLoc []*Coordinate `json:"drop_off_loc"`
	Date string `json:"date"`
	NoOfPassangers int `json:"no_off_passangers"`
	Price int64 `json:"price"`
	VehicleType string `json:"vehicle_type"`
}