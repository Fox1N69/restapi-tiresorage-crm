package middlewares

import "github.com/gofiber/fiber/v3"

type AuthMiddleware struct {
}

type AuthMiddlewareI interface {
	LoginMiddleware(c fiber.Ctx) error
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (aw *AuthMiddleware) LoginMiddleware(c fiber.Ctx) error {
	return nil
}
