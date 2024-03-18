package repository

import (
	"crud-crm/pkg/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

type AuthRepositoryI interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

func NewAuthReposit(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (ar *AuthRepository) CreateUser(user *models.User) error {
	return ar.DB.Create(user).Error
}

func (ar *AuthRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User

	if err := ar.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ar *AuthRepository) UpdateUser(user *models.User) error {
	return ar.DB.Save(user).Error
}

func (ar *AuthRepository) DeleteUser(user *models.User) error {
	return ar.DB.Delete(user).Error
}
