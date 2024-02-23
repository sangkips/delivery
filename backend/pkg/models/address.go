package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	City      string `json:"city"`
	Town      string `json:"town"`
	Street    string `json:"street"`
	Shop_name string `json:"shop name"`
}
