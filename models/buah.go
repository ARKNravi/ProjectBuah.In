package models

import (
	"gorm.io/gorm"
)

type Buah struct {
	gorm.Model
	Nama        string `json:"nama" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint   `json:"price" binding:"required,number"`
	Discount    uint   `json:"discount" binding:"required,number"`
	Quantity    uint   `json:"quantity" binding:"required"`
}

func (Buah) TableName() string {
	return "buah"
}
