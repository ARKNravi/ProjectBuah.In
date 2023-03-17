package repository

import (
	"ProjectBuahIn/models"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepository interface {
	OrderProduct(int, int, int) error
}

type orderRepository struct {
	connection *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		connection: DB(),
	}
}

func (db *orderRepository) OrderProduct(userID int, buahID int, quantity int) error {
	var buah models.Buah
	if err := db.connection.First(&buah, buahID).Error; err != nil {
		return err
	}

	// Check if there is enough stock
	if int(buah.Stok) < quantity {
		return fmt.Errorf("Not enough stock for %s. Available stock: %d", buah.Nama, buah.Stok)
	}

	totalPrice := uint(quantity) * buah.Price

	buah.Stok = buah.Stok - uint(quantity)
	if err := db.connection.Save(&buah).Error; err != nil {
		return err
	}

	return db.connection.Create(&models.Order{
		BuahID:     uint(buahID),
		UserID:     uint(userID),
		Quantity:   uint(quantity),
		Totalprice: totalPrice,
	}).Error
}
