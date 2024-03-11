package handlers

import (
	"crud-crm/pkg/repository"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetAllClients(c fiber.Ctx) error {
	return nil
}

func (h *Handler) GetClientByID(c fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.mainRepo.Client.GetClientByID(uint(clientID))
	if err != nil {
		return err
	}
	return c.JSON(user)
}
