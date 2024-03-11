package routers

import (
	"crud-crm/pkg/handlers"

	"github.com/gofiber/fiber/v3"
)

type Router struct {
	handler handlers.Handler
}

func NewRouter(h handlers.Handler) *Router {
	return &Router{handler: h}
}

func (r *Router) RouterSetup(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Get("/clients", r.handler.GetAllClients)
		api.Get("/get-crequests", r.handler.GetAllCrequests)
	}

}
