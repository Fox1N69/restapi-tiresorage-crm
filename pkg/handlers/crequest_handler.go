package handlers

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
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

func (h *Handler) GetCrequestByBranch(c fiber.Ctx) error {
	branch := c.Params("branch")

	crequest, err := h.repository.Crequest.GetCrequestByBranch(branch)
	if err != nil {
		return err
	}

	return c.JSON(crequest)
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

	id, err := stringToUint(idParam)
	if err != nil {
		return c.JSON(400, "Invalid ID")
	}

	crequest, err := h.repository.Crequest.GetCrequestByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось получить Crequest"})
	}

	var updateCrequest models.Crequests
	if err := json.Unmarshal(c.Body(), &updateCrequest); err != nil {
		return err
	}

	if updateCrequest.ID != 0 {
		crequest.ID = updateCrequest.ID
	}

	if updateCrequest.FIO != "" {
		crequest.FIO = updateCrequest.FIO
	}

	if updateCrequest.StorageCell != "" {
		crequest.StorageCell = updateCrequest.StorageCell
	}

	if updateCrequest.DiskSize != "" {
		crequest.DiskSize = updateCrequest.DiskSize
	}

	if updateCrequest.Price != "" {
		crequest.Price = updateCrequest.Price
	}

	if updateCrequest.IsPaid != "" {
		crequest.IsPaid = updateCrequest.IsPaid
	}

	if updateCrequest.Status != "" {
		crequest.Status = updateCrequest.Status
	}

	if err := h.repository.Crequest.UpdateCrequest(*crequest); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to update crequest")
	}

	return c.JSON(crequest)
}

func (h *Handler) DeleteCrequest(c fiber.Ctx) error {
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

	if err := h.repository.Crequest.DeleteCrequest(*crequest); err != nil {
		return c.JSON(400, "Error delete crequest")
	}

	return c.JSON(fiber.Map{
		"message": "delete successfully",
	})
}

func (h *Handler) DeleteWithClient(c fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id"})
	}

	crequestID, err := stringToUint(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := h.repository.Crequest.DeleteCrequestByID(crequestID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Crequest with ID %s deleted", id)})
}
