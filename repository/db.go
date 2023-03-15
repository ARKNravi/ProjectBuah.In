package repository

import (
	"ProjectBuahIn/models"
	"log"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func DB() *gorm.DB {

	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Error")
	}

	db.AutoMigrate(&models.User{}, &models.Buah{}, &models.Order{}, &models.Cart{}, &models.Address{})
	return db

}
