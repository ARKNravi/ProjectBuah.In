package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	NoTelp   string `json:"notelp" binding:"required"`
}

func (User) TableName() string {
	return "users"
}
