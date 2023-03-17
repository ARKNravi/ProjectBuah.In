package repository

import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserRepository -> User CRUD
type UserRepository interface {
	AddUser(models.User) (models.User, error)
	GetUser(int) (models.User, error)
	GetByUsername(string) (models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(models.User) (models.User, error)
	GetProductOrdered(int) ([]models.Order, error)
}

type userRepository struct {
	connection *gorm.DB
}

// NewUserRepository --> returns new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}

func (db *userRepository) GetUser(id int) (user models.User, err error) {
	return user, db.connection.First(&user, id).Error
}

func (db *userRepository) GetByUsername(username string) (user models.User, err error) {
	return user, db.connection.First(&user, "username=?", username).Error
}

func (db *userRepository) GetAllUser() (users []models.User, err error) {
	return users, db.connection.Find(&users).Error
}

func (db *userRepository) AddUser(user models.User) (models.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) UpdateUser(user models.User) (models.User, error) {
	err := db.connection.Model(&models.User{}).Where("id=?", user.ID).Updates(&user)
	if err.Error != nil {
		return models.User{}, err.Error
	}
	return user, nil
}

func (db *userRepository) DeleteUser(user models.User) (models.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetProductOrdered(userID int) (orders []models.Order, err error) {
	return orders, db.connection.Preload(clause.Associations).Where("user_id = ?", userID).Find(&orders).Error
}
