package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     int16       `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	Order_Cart ProductUser `json:"order" gorm:"foreignKey:UserID;references:ID"`
	Price      int         `json:"price"`
	Discount   *int        `json:"discount"`
}
