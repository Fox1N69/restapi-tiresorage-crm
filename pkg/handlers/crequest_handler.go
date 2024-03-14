package handlers

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/models"
	"encoding/json"
	"errors"

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
	data := new(models.Crequests)

	if err := json.Unmarshal(c.Body(), data); err != nil {
		return err
	}

	if err := h.mainRepo.Crequest.CreateCrequest(data); err != nil {
		return err
	}
	return c.JSON(data)
}

func (h *Handler) GetCrequestByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errors.New("ID is empty")
	}

	creqID, err := h.ConvertUint(id)
	if err != nil {
		return err
	}

	crequest, err := h.mainRepo.Crequest.GetCrequestByID(creqID)
	if err != nil {
		return err
	}

	return c.JSON(crequest)
}

func (h *Handler) UpdateCrequest(c fiber.Ctx) error {
	return nil
}

func (h *Handler) DeleteCrequest(c fiber.Ctx) error {
	return nil
}
