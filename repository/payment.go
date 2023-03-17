package repository

import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentRepository interface {
	AddPayment(models.Payment, int, int) (models.Payment, error)
	GetPayment(int) (models.Payment, error)
}

type paymentRepository struct {
	connection *gorm.DB
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{
		connection: DB(),
	}
}

func (db *paymentRepository) AddPayment(payment models.Payment, userID int, checkoutID int) (models.Payment, error) {

	var user models.User
	if err := db.connection.First(&user, userID).Error; err != nil {
		return models.Payment{}, err
	}
	var checkout models.Checkout
	if err := db.connection.First(&checkout, checkoutID).Error; err != nil {
		return models.Payment{}, err
	}

	// Set the user ID for the address
	payment.UserID = uint(userID)
	payment.CheckoutID = uint(checkoutID)

	// Save the address to the database
	if err := db.connection.Preload(clause.Associations).Create(&payment).Error; err != nil {
		return models.Payment{
			UserID:     uint(userID),
			CheckoutID: uint(checkoutID),
		}, err
	}

	return payment, nil
}

func (db *paymentRepository) GetPayment(id int) (payment models.Payment, err error) {
	return payment, db.connection.Preload(clause.Associations).First(&payment, id).Error
}
