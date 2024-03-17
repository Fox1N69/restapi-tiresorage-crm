package controllers

import "gorm.io/gorm"


type Controllers struct {
	authCont *AuthController
}

func NewControllers(db *gorm.DB) *Controllers {
	return &Controllers{authCont: NewAuthController(db)}
}
