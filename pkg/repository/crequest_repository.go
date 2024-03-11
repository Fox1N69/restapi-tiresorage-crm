package repository

import "gorm.io/gorm"

type CrequestRepository struct {
	db *gorm.DB
}

func NewCrequestRepository(db *gorm.DB) *CrequestRepository {
	return &CrequestRepository{}
}