package controllers

import "gorm.io/gorm"

type Controllers struct {
	Auth *AuthController
}

func NewControllers(db *gorm.DB) *Controllers {
	return &Controllers{Auth: NewAuthController(db)}
}
