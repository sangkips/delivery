package models

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	UserID        string
	User          *User   `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
	PaymentID     int     `json:"-"`
	PaymentMethod string  `json:"payment_method"` // mpesa, cash, card, credit card
	Amount        float32 `json:"amount"`
}
