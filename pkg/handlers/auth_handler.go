package handlers

import (

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Register(c fiber.Ctx) error {
	return h.controller.Auth.Register(c)
}

func (h *Handler) Login(c fiber.Ctx) error {
	return h.controller.Auth.Login(c)
}

func (h *Handler) Logout(c fiber.Ctx) error {
	return h.controller.Auth.Logout(c)
}

func (h *Handler) GetUsers(c fiber.Ctx) error {
	return h.controller.Auth.GetUsers(c)
}
