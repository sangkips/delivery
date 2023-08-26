package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"size:100;not null;" json:"name"`
	Price       float32 `gorm:"size:100;not null;" json:"price"`
	Description string  `json:"description"`
	Brand       string  `gorm:"size:100;not null;" json:"brand"`
	Image       string  `json:"image"`
}
