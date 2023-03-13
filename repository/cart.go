package repository

import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
)

// ProductRepository --> Interface to ProductRepository
type CartRepository interface {
	GetCart(int) (models.Cart, error)
	GetAllCart() ([]models.Cart, error)
	AddCart(int, int) error
	UpdateCart(int, int, int, models.Cart) error
	DeleteCart(int, int, int, models.Cart) error
}

type cartRepository struct {
	connection *gorm.DB
}

// NewProductRepository --> returns new product repository
func NewCartRepository() CartRepository {
	return &cartRepository{
		connection: DB(),
	}
}

func (db *cartRepository) GetCart(id int) (cart models.Cart, err error) {
	return cart, db.connection.First(&cart, id).Error
}

func (db *cartRepository) GetAllCart() (carts []models.Cart, err error) {
	return carts, db.connection.Find(&carts).Error
}

func (db *cartRepository) AddCart(userID int, productID int) error {
	return db.connection.Create(&models.Cart{
		ProductID: uint(productID),
		UserID:    uint(userID),
	}).Error
}

func (db *cartRepository) UpdateCart(userID int, productID int, id int, cart models.Cart) error {
	if err := db.connection.First(&cart, cart.ID).Error; err != nil {
		return err
	}
	return db.connection.Create(&models.Cart{
		ProductID: uint(productID),
		UserID:    uint(userID),
	}).Error
}

func (db *cartRepository) DeleteCart(userID int, productID int, id int, cart models.Cart) error {
	if err := db.connection.First(&cart, cart.ID).Error; err != nil {
		return err
	}
	return db.connection.Delete(&models.Cart{
		ProductID: uint(productID),
		UserID:    uint(userID),
	}).Error
}
