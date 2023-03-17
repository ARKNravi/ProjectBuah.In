package repository

//repository
import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AddressRepository interface {
	GetAddress(int) (models.Address, error)
	GetAllAddress(int) ([]models.Address, error)
	AddAddress(models.Address, int) (models.Address, error)
	UpdateAddress(models.Address) (models.Address, error)
	DeleteAddress(models.Address) (models.Address, error)
}

type addressRepository struct {
	connection *gorm.DB
}

func NewAddressRepository() AddressRepository {
	return &addressRepository{
		connection: DB(),
	}
}

func (db *addressRepository) GetAddress(id int) (address models.Address, err error) {
	return address, db.connection.Preload(clause.Associations).First(&address, id).Error
}

func (db *addressRepository) GetAllAddress(userID int) (addresss []models.Address, err error) {
	return addresss, db.connection.Preload(clause.Associations).Where("user_id = ?", userID).Find(&addresss).Error
}

func (db *addressRepository) AddAddress(address models.Address, userID int) (models.Address, error) {
	var user models.User
	if err := db.connection.First(&user, userID).Error; err != nil {
		return models.Address{}, err
	}

	// Set the user ID for the address
	address.UserID = uint(userID)

	// Save the address to the database
	if err := db.connection.Preload(clause.Associations).Create(&address).Error; err != nil {
		return models.Address{
			UserID: uint(userID),
		}, err
	}

	return address, nil
}

func (db *addressRepository) UpdateAddress(address models.Address) (models.Address, error) {
	err := db.connection.Model(&models.User{}).Where("id=?", address.ID).Updates(&address)
	if err.Error != nil {
		return models.Address{}, err.Error
	}
	return address, nil
}

func (db *addressRepository) DeleteAddress(address models.Address) (models.Address, error) {
	if err := db.connection.First(&address, address.ID).Error; err != nil {
		return address, err
	}
	return address, db.connection.Delete(&address).Error
}
