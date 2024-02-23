package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name   string        `json:"first_name" validate:"required, min=2, max=150"`
	Last_name    string        `json:"last_name" validate:"required, min=2, max=150"`
	Email        string        `gorm:"unique" json:"email"`
	Password     string        `json:"password" validate:"required, min=6"`
	Phone        string        `gorm:"unique" json:"phone number" validate:"required, min=10, max=15"`
	UserID       string        `json:"-"`
	User_Order   []ProductUser `json:"user_order" gorm:"foreignKey:UserID;references:ID"`
	Address      []Address     `json:"address" gorm:"foreignKey:UserID;references:ID"`
	Order_Status []Order       `json:"order_status" gorm:"foreignKey:UserID;references:ID"`
}
