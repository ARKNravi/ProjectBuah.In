package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	User       User     `gorm:"foreignkey:UserID"`
	Checkout   Checkout `gorm:"foreignkey:CheckoutID"`
	Imagelink  string   `json:"imagelink"`
	UserID     uint
	CheckoutID uint
}

func (Payment) TableName() string {
	return "payments"
}
