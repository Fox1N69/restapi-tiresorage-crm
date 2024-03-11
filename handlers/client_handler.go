package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Test(c fiber.Ctx) error {
	return c.SendString("Hello fiber")
}
