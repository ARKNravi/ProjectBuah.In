package initializer

import (
	"ProjectBuahIn/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB 2345")
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Buah{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Order{})

}
