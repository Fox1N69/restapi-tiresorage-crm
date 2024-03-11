package repository

import "gorm.io/gorm"

type MainRepository struct {
	Client   *ClientRepository
	Crequest *CrequestRepository
}

func NewMainRepository(db *gorm.DB) *MainRepository {
	return &MainRepository{
		Client:   NewClientRepository(db),
		Crequest: NewCrequestRepository(db),
	}
}
