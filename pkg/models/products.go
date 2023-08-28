package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;" json:"name"`
	Price       uint64 `gorm:"size:100;not null;" json:"price"`
	Description string `json:"description"`
	Brand       string `gorm:"size:100;not null;" json:"brand"`
	Image       string `json:"image"`
}

type ProductUser struct {
	gorm.Model
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Brand       string `json:"brand"`
	Image       string `json:"image"`
	UserID      string `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
}
