package repository

import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CheckoutRepository interface {
	GetCheckout(int) (models.Checkout, error)
	AddCheckout(int, int, int) error
}

type checkoutRepository struct {
	connection *gorm.DB
}

// NewProductRepository --> returns new product repository
func NewCheckoutRepository() CheckoutRepository {
	return &checkoutRepository{
		connection: DB(),
	}
}

func (db *checkoutRepository) GetCheckout(id int) (checkout models.Checkout, err error) {
	return checkout, db.connection.Preload(clause.Associations).First(&checkout, id).Error
}

func (db *checkoutRepository) AddCheckout(userID int, addressID int, cartID int) error {
	return db.connection.Create(&models.Checkout{
		UserID:    uint(userID),
		AddressID: uint(addressID),
		CartID:    uint(cartID),
	}).Error

}
