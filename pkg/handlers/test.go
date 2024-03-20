package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Test(c fiber.Ctx) error {
	data := struct {
		Message string `json:"message"`
	}{}

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": data.Message,
	})
}
