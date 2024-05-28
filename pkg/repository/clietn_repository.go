package repository

import (
	"crud-crm/pkg/models"
	"errors"

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
	if r == nil {
		return errors.New("ClientRepository is null")
	}
	return r.db.Create(client).Error
}

func (r *ClientRepository) GetClientByID(id uint) (*models.Client, error) {
	var client models.Client
	if err := r.db.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientRepository) GetClientByBranch(branch string) ([]models.Client, error) {
	var client []models.Client

	if err := r.db.Where("branch = ?", branch).Find(&client).Error; err != nil {
		return nil, err
	}

	return client, nil
}

func (r *ClientRepository) UpdateClient(client *models.Client) error {
	return r.db.Save(&client).Error
}

func (r *ClientRepository) DeleteClient(client *models.Client) error {
	if err := r.db.Exec("ALTER SEQUENCE clients_id_seq RESTART WITH 1").Error; err != nil {
		r.db.Rollback()
		return err
	}

	return r.db.Delete(&client).Error
}
