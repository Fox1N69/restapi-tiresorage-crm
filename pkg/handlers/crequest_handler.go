package handlers

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/models"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetAllCrequests(c fiber.Ctx) error {
	var data models.Crequests
	if err := database.DB.Find(&data).Error; err != nil {
		return err
	}

	return c.JSON(data)
}

func (h *Handler) CreateCrequest(c fiber.Ctx) error {
	return nil
}

func (h *Handler) GetCrequestByID(c fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdateCrequest(c fiber.Ctx) error {
	return nil
}

func (h *Handler) DeleteCrequest(c fiber.Ctx) error {
	return nil
}
