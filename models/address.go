package models

//models
import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	User         User `gorm:"foreignkey:UserID"`
	UserID       uint
	Namapenerima *string `json:"namapenerima"`
	Nomorhp      *string `json:"nomorhp"`
	Label        *string `json:"label"`
	Kota         *string `json:"kota"`
	Alamat       *string `json:"alamat"`
}

func (Address) TableName() string {
	return "address"
}
