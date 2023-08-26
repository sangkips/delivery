package models

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	Amount    float32 `json:"amount"`
	ProductID int     `json:"-"`
	Product   Product `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}
