package handlers

import (
	"crud-crm/pkg/models"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (h *Handler) Register(c fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return err
	}

	if err := h.mainRepo.Auth.CreateUser(&user); err != nil {
		return err
	}

	return nil
}

func (h *Handler) Login(c fiber.Ctx) error {
	return nil
}

func (h *Handler) Logout(c fiber.Ctx) error {
	return nil
}

func (h *Handler) GetUsers(c fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.JSON(fiber.Map{
			err.Error(): "error unmarshl body",
		})
	}

	return c.JSON(&user)
}
