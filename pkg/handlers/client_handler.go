package handlers

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/models"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetAllClients(c fiber.Ctx) error {
	var clients []models.Client
	database.DB.Find(&clients)

	return c.JSON(clients)
}


func (h *Handler) CreateClient(c fiber.Ctx) error {
	client := new(models.Client)
	if err := json.Unmarshal(c.Body(), &client); err != nil {
		return err
	}

	if err := h.repository.Client.CreateClient(client); err != nil {
		panic(err)
	}

	return c.JSON(client)
}

func (h *Handler) GetClientByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(fiber.Map{
			"message": "ID is empty",
		})
	}

	clientID, err := h.ConvertUint(id)
	if err != nil {
		return err
	}

	user, err := h.repository.Client.GetClientByID(clientID)
	if err != nil {
		panic(err)
	}

	return c.JSON(user)
}

func (h *Handler) UpdateClient(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(fiber.Map{
			"message": "ID empty",
		})
	}

	clientID, err := h.ConvertUint(id)
	if err != nil {
		return err
	}

	exitClient, err := h.repository.Client.GetClientByID(clientID)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(c.Body(), &exitClient); err != nil {
		return err
	}

	if err := h.repository.Client.UpdateClient(exitClient); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "client seccessfully change",
	})
}

func (h *Handler) DeleteClient(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(fiber.Map{
			"message": "ID empty",
		})
	}
	clientID, err := h.ConvertUint(id)
	if err != nil {
		return err
	}

	exitClient, err := h.repository.Client.GetClientByID(clientID)
	if err != nil {
		return err
	}

	if err := h.repository.Client.DeleteClient(exitClient); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Client seccessfully deleted",
	})
}
