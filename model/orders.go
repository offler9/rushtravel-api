package model

import (
	"github.com/jinzhu/gorm"
	//"github.com/lib/pq"
)

type Order struct {
	gorm.Model
	OrderId string `gorm:"column:order_id" json:"orderId"`
	Username string `gorm:"column:username" json:"username"`
	PickUpLoc []string `gorm:"type:[]string;column:pick_up_loc"`
	DropOffLoc []string `gorm:"type:[]string;column:drop_off_loc"`
	Date string `gorm:"column:date" json:"date"`
	Days string `gorm:"column:days" json:"days"`
	NoOfPassangers int `gorm:"column:no_off_passangers" json:"no_off_passangers"`
	Price int `gorm:"column:price" json:"price"`
	VehicleType string `gorm:"column:vehicle_type" json:"vehicle_type"`
	Status string `gorm:"column:status" json:"status"`
	Note string `gorm:"column:note" json:"note"`
}