package repository

import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
)

type AddressRepository interface {
	GetAddress(int) (models.Address, error)
	GetAllAddress() ([]models.Address, error)
	AddAddress(models.Address) (models.Address, error)
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
	return address, db.connection.First(&address, id).Error
}

func (db *addressRepository) GetAllAddress() (addresss []models.Address, err error) {
	return addresss, db.connection.Find(&addresss).Error
}

func (db *addressRepository) AddAddress(address models.Address) (models.Address, error) {
	return address, db.connection.Create(&address).Error
}

func (db *addressRepository) UpdateAddress(address models.Address) (models.Address, error) {
	if err := db.connection.First(&address, address.ID).Error; err != nil {
		return address, err
	}
	return address, db.connection.Model(&address).Updates(&address).Error
}

func (db *addressRepository) DeleteAddress(address models.Address) (models.Address, error) {
	if err := db.connection.First(&address, address.ID).Error; err != nil {
		return address, err
	}
	return address, db.connection.Delete(&address).Error
}
