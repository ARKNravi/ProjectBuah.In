package repository

//repository
import (
	"ProjectBuahIn/models"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProductRepository --> Interface to ProductRepository
type CartRepository interface {
	GetCart(int) (models.Cart, error)
	GetAllCart(int) ([]models.Cart, error)
	AddCart(int, int, int) error
	UpdateCart(models.Cart) (models.Cart, error)
	DeleteCart(models.Cart) (models.Cart, error)
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
	return cart, db.connection.Preload(clause.Associations).First(&cart, id).Error
}

func (db *cartRepository) AddCart(userID int, buahID int, quantity int) error {
	var buah models.Buah
	if err := db.connection.First(&buah, buahID).Error; err != nil {
		return err
	}

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

func (db *cartRepository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := db.connection.Model(&models.Cart{}).Where("id=?", cart.ID).Updates(&cart)
	if err.Error != nil {
		return models.Cart{}, err.Error
	}
	return cart, nil
}

func (db *cartRepository) DeleteCart(cart models.Cart) (models.Cart, error) {
	if err := db.connection.First(&cart, cart.ID).Error; err != nil {
		return cart, err
	}
	return cart, db.connection.Delete(&cart).Error
}

func (db *cartRepository) GetAllCart(userID int) (carts []models.Cart, err error) {
	return carts, db.connection.Preload(clause.Associations).Where("user_id = ?", userID).Find(&carts).Error
}
