package models

//models
import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password" `
	NoTelp   string `json:"notelp"`
}

func (User) TableName() string {
	return "users"
}
