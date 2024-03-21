package controllers

import (
	"crud-crm/pkg/models"
	"crud-crm/pkg/repository"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthController struct {
	DB         *gorm.DB
	repository repository.AuthRepository
}

type AuthControllerI interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
	Logout(c fiber.Ctx) error
	GetUser(c fiber.Ctx) error
	GenerateToken(user *models.User) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db, repository: repository.AuthRepository{}}
}

const (
	secretKey = "secret"
)

func (ac *AuthController) Login(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Register(c fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return err
	}

	if err := ac.repository.CreateUser(&user); err != nil {
		return err
	}

	token, err := ac.GenerateToken(&user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "user create",
		"token":   token,
	})
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) GetUsers(c fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.JSON(fiber.Map{
			err.Error(): "error unmarshl body",
		})
	}

	return c.JSON(&user)
}

// Generate jwt token
func (ac *AuthController) GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID
	claims["username"] = user.Username

	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ac *AuthController) VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
}
