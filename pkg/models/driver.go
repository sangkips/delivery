package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	Name         string `gorm:"size:100;not null;" json:"name"`
	Registration string `gorm:"size:100;not null;" json:"registration"`
	IdentityCard string `gorm:"size:10;not null;" json:"id"`
}
