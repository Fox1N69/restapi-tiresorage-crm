package middlewares

import (
	"crud-crm/pkg/models"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type AuthMiddleware struct {
}

type AuthMiddlewareI interface {
	LoginMiddleware(c fiber.Ctx) error
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (aw *AuthMiddleware) LoginMiddleware(c fiber.Ctx) error {
	var login models.User

	if err := json.Unmarshal(c.Body(), &login); err != nil {
		return c.JSON(fiber.Map{
			"message": "Invalid input data",
		})
	}

	if login.Password == nil || login.Username == "" {
		return c.JSON(fiber.Map{
			"message": "Empty username or password",
		})
	}

	return c.Next()
}
