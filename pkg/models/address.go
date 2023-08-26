package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID      string `json:"-"`
	User        *User  `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
}
