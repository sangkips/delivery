package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name   string        `json:"first_name"`
	Last_name    string        `json:"last_name"`
	Email        string        `gorm:"unique" json:"email"`
	Password     string        `json:"password"`
	Phone        string        `json:"phone number"`
	UserID       string        `json:"-"`
	User_Order   []ProductUser `json:"user_order" gorm:"foreignKey:UserID;references:ID"`
	Address      []Address     `json:"address" gorm:"foreignKey:UserID;references:ID"`
	Order_Status []Order       `json:"order_status" gorm:"foreignKey:UserID;references:ID"`
}
