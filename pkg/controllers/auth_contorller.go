package controllers

import (
	"crud-crm/pkg/models"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

type AuthControllerI interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
	Logout(c fiber.Ctx) error
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}


func (ac *AuthController) Login(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Register(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	return nil
}
