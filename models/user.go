package models

//models
import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	NoTelp   string `json:"notelp"`
}

func (User) TableName() string {
	return "users"
}
