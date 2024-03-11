package handlers

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/models"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetAllClients(c fiber.Ctx) error {
	return nil
}

func (h *Handler) CreateClient(c fiber.Ctx) error {
	client := new(models.Client)
	if err := json.Unmarshal(c.Body(), &client); err != nil {
		return err
	}

	if err := h.mainRepo.Client.CreateClient(client); err != nil {
		panic(err)
	}

	return c.JSON(client)
}

func (h *Handler) GetClientByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(fiber.Map{
			"message": "ID empty",
		})
	}
	clientID := ConverUint(id)

	user, err := h.mainRepo.Client.GetClientByID(clientID)
	if err != nil {
		panic(err)
	}

	return c.JSON(user)
}

func (h *Handler) UpdateClient(c fiber.Ctx) error {
	var client models.Client
	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(fiber.Map{
			"message": "ID empty",
		})
	}

	if err := database.DB.Where("id = ?").First(&client).Error; err != nil {
		return err
	}

	if err := json.Unmarshal(c.Body(), &client); err != nil {
		return err
	}

	if err := h.mainRepo.Client.UpdateClient(&client); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "client seccessfully change",
	})
}

func (h *Handler) DeleteClient(c fiber.Ctx) error {
	return nil
}
