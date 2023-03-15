package repository

import (
	"ProjectBuahIn/models"

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
	return db.connection.Create(&models.Order{
		BuahID:   uint(buahID),
		UserID:   uint(userID),
		Quantity: quantity,
	}).Error

}
