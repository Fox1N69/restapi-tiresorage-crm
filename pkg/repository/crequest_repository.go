package repository

import (
	"crud-crm/pkg/models"

	"gorm.io/gorm"
)

type CrequestRepository struct {
	db *gorm.DB
}

type CrequestRepositoryI interface {
	CreateCrequest(crequest models.Crequests) error
	GetCrequestByID(id uint) (*models.Crequests, error)
	UpdateCrequest(crequest models.Crequests) error
	DeleteCrequest(crequest models.Crequests) error
}

func NewCrequestRepository(db *gorm.DB) *CrequestRepository {
	return &CrequestRepository{db: db}
}

func (cr *CrequestRepository) CreateCrequest(crequest models.Crequests) error {
	return nil
}

func (cr *CrequestRepository) GetCrequestByID(id uint) (*models.Crequests, error) {
	return nil, nil
}

func (cr *CrequestRepository) UpdateCrequest(crequest models.Crequests) error {
	return nil
}

func (cr *CrequestRepository) DeleteCrequest(crequest models.Crequests) error {
	return nil
}
