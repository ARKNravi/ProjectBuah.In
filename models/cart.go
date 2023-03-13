package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User       User `gorm:"foreignkey:UserID"`
	Product    Buah `gorm:"foreignkey:ProductID"`
	UserID     uint
	ProductID  uint
	Quantity   uint `json:"quantity"`
	TotalPrice uint `json:"totalprice"`
}

func (Cart) TableName() string {
	return "cart"
}
