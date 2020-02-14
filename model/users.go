package model

import (
	"github.com/jinzhu/gorm"
)

// User is a model for user table
type User struct {
	gorm.Model
	// Email    string `json:"email"`
	// Name     string `json:"name"`
	Username string `gorm:"primary_key;not null"`
	PhoneNum int `json:"phone_number"`
	Password string `json:"password"`
}
