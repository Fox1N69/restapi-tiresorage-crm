package repository

import (
	"crud-crm/pkg/models"

	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

type ClientRepositoryI interface {
	CreateClient(client *models.Client) error
	GetClientByID(id uint) (*models.Client, error)
	UpdateClient(client *models.Client) error
	DeleteClient(client *models.Client) error
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) CreateClient(client *models.Client) error {
	return r.db.Create(client).Error
}

func (r *ClientRepository) GetClientByID(id uint) (*models.Client, error) {
	var client models.Client
	if err := r.db.First(&client, id).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientRepository) UpdateClient(client *models.Client) error {
	return r.db.Save(&client).Error
}

func (r *ClientRepository) DeleteClient(client *models.Client) error {
	return r.db.Delete(&client).Error
}
