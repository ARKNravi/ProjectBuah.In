package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User     User `gorm:"foreignkey:UserID"`
	Buah     Buah `gorm:"foreignkey:BuahID"`
	UserID   uint
	BuahID   uint
	Quantity int  `json:"quantity"`
	Status   bool `json:"status"`
}

func (Order) TableName() string {
	return "orders"
}
