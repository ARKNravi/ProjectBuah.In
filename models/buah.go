package models

//models
import (
	"gorm.io/gorm"
)

type Buah struct {
	gorm.Model
	Nama        string `json:"nama" binding:"required"`
	Kondisi     string `json:"kondisi" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint   `json:"price" binding:"required,number"`
	Discount    uint   `json:"discount" binding:"required,number"`
	Berat       uint   `json:"berat" binding:"required,number"`
	Stok        uint   `json:"stok" binding:"required"`
	Alamatbuah  string `json:"alamatbuah" binding:"required"`
	Rating      string `json:"rating" binding:"required"`
	Sold        string `json:"sold" binding:"required"`
	ImageLink   string `json:"image_link"`
}

func (Buah) TableName() string {
	return "buah"
}
