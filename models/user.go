package models

//models
import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique"`
	Email    string `json:"email" `
	Password string `json:"password" binding:"required"`
	NoTelp   string `json:"notelp"`
}

func (User) TableName() string {
	return "users"
}
