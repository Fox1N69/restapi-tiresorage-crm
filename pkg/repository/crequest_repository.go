package repository

import (
	"crud-crm/pkg/models"

	"gorm.io/gorm"
)

type CrequestRepositoryI interface {
	CreateCrequst(crequest models.Crequests) error
}
type CrequestRepository struct {
	db *gorm.DB
}

func NewCrequestRepository(db *gorm.DB) *CrequestRepository {
	return &CrequestRepository{}
}
