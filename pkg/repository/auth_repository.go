package repository

import "gorm.io/gorm"

type AuthRepository struct {
	DB *gorm.DB
}

type AuthRepositoryI interface {
}

func NewAuthReposit(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}
