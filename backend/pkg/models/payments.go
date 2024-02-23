package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID         string `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	PaymentID      int    `json:"-"`
	Payment_Method string `json:"payment_method"` // mpesa, cash, card, credit card
	Amount         uint64 `json:"amount"`
}
