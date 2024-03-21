package handlers

import (
	"crud-crm/pkg/models"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Register(c fiber.Ctx) error {
	data := new(models.User)

	if err := json.Unmarshal(c.Body(), data); err != nil {
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
	return nil
}
