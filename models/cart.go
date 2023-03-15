package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User       User `gorm:"foreignkey:UserID"`
	Buah       Buah `gorm:"foreignkey:BuahID"`
	UserID     uint
	BuahID     uint
	Quantity   uint `json:"quantity"`
	Totalprice uint `json:"totalprice"`
}

func (Cart) TableName() string {
	return "cart"
}
