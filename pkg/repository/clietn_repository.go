package repository

import (
	"crud-crm/pkg/models"

	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{}
}

func (r *ClientRepository) CreateClient(client *models.Clients) error {
	return r.db.Create(&client).Error
}

func (r *ClientRepository) GetClientByID(id uint) (*models.Clients, error) {
	var client models.Clients
	if err := r.db.First(&client, id).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientRepository) UpdateClient(client *models.Clients) error {
	return r.db.Save(&client).Error
}

func (r *ClientRepository) DeleteClient(client *models.Clients) error {
	return r.db.Delete(&client).Error
}
