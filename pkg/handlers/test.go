package handlers

import (
	"crud-crm/pkg/models"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Test(c fiber.Ctx) error {
	var data models.Client

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return err
	}

	h.repository.Client.CreateClient(&data)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": data,
	})
}
