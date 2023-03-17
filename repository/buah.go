package repository

//repository
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
	GetBuahByKondisi(string) ([]models.Buah, error)
	GetBuahByPriceDescending() ([]models.Buah, error)
	GetBuahByPriceAscending() ([]models.Buah, error)
	SeedData() error
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

func (db *buahRepository) GetBuahByKondisi(kondisi string) ([]models.Buah, error) {
	var buahs []models.Buah
	err := db.connection.Where("kondisi = ?", kondisi).Find(&buahs).Error
	if err != nil {
		return nil, err
	}
	return buahs, nil
}

func (db *buahRepository) GetBuahByPriceDescending() ([]models.Buah, error) {
	var buahs []models.Buah
	err := db.connection.Order("price desc").Find(&buahs).Error
	if err != nil {
		return nil, err
	}
	return buahs, nil
}

func (db *buahRepository) GetBuahByPriceAscending() ([]models.Buah, error) {
	var buahs []models.Buah
	err := db.connection.Order("price asc").Find(&buahs).Error
	if err != nil {
		return nil, err
	}
	return buahs, nil
}

func (db *buahRepository) SeedData() error {
	buahs := []models.Buah{
		{Nama: "Apel", Kondisi: "Segar", Description: "Buah apel yang segar dan enak", Price: 5000, Discount: 10, Berat: 100, Stok: 50, Alamatbuah: "Jalan Apel No. 1", ImageLink: "https://example.com/apel.jpg"},
		{Nama: "Mangga", Kondisi: "Matang", Description: "Buah mangga yang matang dan manis", Price: 10000, Discount: 0, Berat: 200, Stok: 30, Alamatbuah: "Jalan Mangga No. 2", ImageLink: "https://example.com/mangga.jpg"},
		{Nama: "Jeruk", Kondisi: "Kecut", Description: "Buah jeruk yang kecut dan segar", Price: 7500, Discount: 5, Berat: 150, Stok: 20, Alamatbuah: "Jalan Jeruk No. 3", ImageLink: "https://example.com/jeruk.jpg"},
	}
	for _, buah := range buahs {
		if err := db.connection.Create(&buah).Error; err != nil {
			return err
		}
	}
	return nil
}
