package handlers

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func stringToUint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(num), nil
}

func (h *Handler) GetAllCrequests(c fiber.Ctx) error {
	var data []models.Crequests

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

	if err := h.repository.Crequest.CreateCrequest(data); err != nil {
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

	crequest, err := h.repository.Crequest.GetCrequestByID(creqID)
	if err != nil {
		return err
	}

	return c.JSON(crequest)
}

func (h *Handler) UpdateCrequest(c fiber.Ctx) error {
	idParam := c.Params("id")

	//convert param id to uint
	id, err := stringToUint(idParam)
	if err != nil {
		return c.JSON(400, "Invalid ID")
	}

	crequest, err := h.repository.Crequest.GetCrequestByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось получить Crequest"})
	}

	var updateCrequest models.Crequests
	if err := json.Unmarshal(c.Body(), &crequest); err != nil {
		return err
	}

	// Обновление полей на основе входящих данных
	if updateCrequest.ID != 0 {
		crequest.ID = updateCrequest.ID
	}

	if updateCrequest.FIO != "" {
		crequest.FIO = updateCrequest.FIO
	}
	
	if updateCrequest.Data != "" {
		crequest.Data = updateCrequest.Data
	}

	if updateCrequest.Price != "" {
		crequest.Price = updateCrequest.Price
	}

	if err := h.service.UpdatePost(id, crequest); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to update post")
	}

	return c.JSON(crequest)
}

func (h *Handler) DeleteCrequest(c fiber.Ctx) error {
	return nil
}
