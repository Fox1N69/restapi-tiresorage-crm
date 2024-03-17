package controllers

import (
	"crud-crm/pkg/database"
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

func (ac *AuthController) GenerateToken(c fiber.Ctx) error {
	var user models.User
	if err := database.DB.Where("username = ?", c.FormValue("username")).First(&user).Error; err != nil {
		return err
	}
	if user.Password != c.FormValue("password") {
		return fiber.ErrUnauthorized
	}


	return c.JSON(fiber.Map{"token": ""})
}

func (ac *AuthController) issueToken(userID uint) (string, error) {
	token, err := ac.GenerateToken(userID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ac *AuthController) Login(c fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return err
	}

	err := ac.DB.Raw("SELECT password FROM users WHERE username = ? LIMIT 1", user.Username).Row().Scan(&user.Password)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, user.Password); err != nil {
		return err
	}

	token, err := ac.issueToken(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token})
}

func (ac *AuthController) Register(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	return nil
}
