package repository

import (
	"gorm.io/gorm"
)

// Создаем тип MainRepository с полями существующих рпепозиториев
type MainRepository struct {
	Client   *ClientRepository
	Crequest *CrequestRepository
}

// Создаем ноавый экземпляр MainRepository и добавляем в него вре репозитории
func NewMainRepository(db *gorm.DB) *MainRepository {
	return &MainRepository{
		Client:   NewClientRepository(db),
		Crequest: NewCrequestRepository(db),
	}
}
