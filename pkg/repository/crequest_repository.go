package repository

import (
	"crud-crm/pkg/models"
	"errors"

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
	GetCrequestByFIO(crequest models.Crequests) (*models.Crequests, error)
}

func NewCrequestRepository(db *gorm.DB) *CrequestRepository {
	return &CrequestRepository{db: db}
}

func (cr *CrequestRepository) CreateCrequest(crequest *models.Crequests) error {
	if cr == nil {
		return errors.New("CrequestRepository is null")
	}
	return cr.db.Create(crequest).Error
}

func (cr *CrequestRepository) GetCrequestByBranch(branch string) ([]models.Crequests, error) {
	var crequest []models.Crequests

	if err := cr.db.Where("branch = ?", branch).Find(&crequest).Error; err != nil {
		return nil, err
	}

	return crequest, nil
}

func (r *CrequestRepository) DeleteCrequestByID(id uint) error {
	if err := r.db.Delete(&models.Crequests{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CrequestRepository) GetCrequestByFIO(fio string) (*models.Crequests, error) {
	var crequest models.Crequests
	result := cr.db.Where("fio = ?", fio).First(&crequest)
	if result.Error != nil {
		return nil, result.Error
	}
	return &crequest, nil
}

func (cr *CrequestRepository) GetCrequestByID(id uint) (*models.Crequests, error) {
	var crequest models.Crequests
	if err := cr.db.Where("id = ?", id).First(&crequest).Error; err != nil {
		return nil, err
	}
	return &crequest, nil
}

func (cr *CrequestRepository) UpdateCrequest(crequest models.Crequests) error {
	return cr.db.Save(&crequest).Error
}

func (cr *CrequestRepository) DeleteCrequest(crequest models.Crequests) error {
	return cr.db.Delete(&crequest).Error
}
