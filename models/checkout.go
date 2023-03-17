package models

import (
	"gorm.io/gorm"
)

type Checkout struct {
	gorm.Model
	User          User    `gorm:"foreignkey:UserID"`
	Address       Address `gorm:"foreignkey:AddressID"`
	Cart          Cart    `gorm:"foreignkey:CartID"`
	UserID        uint
	AddressID     uint
	CartID        uint
	TotalPrice    uint   `json:"total_price"`
	PaymentMethod string `json:"payment_method"`
	Shipping      string `json:"shipping"`
}
