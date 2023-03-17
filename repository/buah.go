package repository

import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
)

type BuahRepository interface {
	GetBuah(int) (models.Buah, error)
	GetAllBuah() ([]models.Buah, error)
	AddBuah(models.Buah) (models.Buah, error)
	UpdateBuah(models.Buah) (models.Buah, error)
	DeleteBuah(models.Buah) (models.Buah, error)
}

type buahRepository struct {
	connection *gorm.DB
}

func NewBuahRepository() BuahRepository {
	return &buahRepository{
		connection: DB(),
	}
}

func (db *buahRepository) GetBuah(id int) (Buah models.Buah, err error) {
	return Buah, db.connection.First(&Buah, id).Error
}

func (db *buahRepository) GetAllBuah() (Buahs []models.Buah, err error) {
	return Buahs, db.connection.Find(&Buahs).Error
}

func (db *buahRepository) AddBuah(Buah models.Buah) (models.Buah, error) {
	return Buah, db.connection.Create(&Buah).Error
}

func (db *buahRepository) UpdateBuah(Buah models.Buah) (models.Buah, error) {
	err := db.connection.Model(&models.User{}).Where("id=?", Buah.ID).Updates(&Buah)
	if err.Error != nil {
		return models.Buah{}, err.Error
	}
	return Buah, nil
}

func (db *buahRepository) DeleteBuah(Buah models.Buah) (models.Buah, error) {
	if err := db.connection.First(&Buah, Buah.ID).Error; err != nil {
		return Buah, err
	}
	return Buah, db.connection.Delete(&Buah).Error
}
