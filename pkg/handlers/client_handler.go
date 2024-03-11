package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetAllClients(c fiber.Ctx) error {
	return nil
}

func (h *Handler) GetClientByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.JSON("ID empty")
	}
	clientID := ConverUint(id)

	user, err := h.mainRepo.Client.GetClientByID(clientID)
	if err != nil {
		panic(err)
	}

	
	return c.JSON(&user)
}
