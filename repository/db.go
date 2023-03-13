package repository

import (
	"ProjectBuahIn/models"
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func DB() *gorm.DB {

	dsn := "root:@tcp(127.0.0.1:3306)/projectbcc5?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Error")
	}

	db.AutoMigrate(&models.User{}, &models.Buah{}, &models.Order{}, &models.Cart{})
	return db

}
