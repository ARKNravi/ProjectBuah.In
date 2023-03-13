package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User      User `gorm:"foreignkey:UserID"`
	Product   Buah `gorm:"foreignkey:ProductID"`
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity"`
}

func (Order) TableName() string {
	return "orders"
}
