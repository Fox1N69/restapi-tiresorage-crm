package routers

import (
	"crud-crm/handlers"

	"github.com/gofiber/fiber/v3"
)

type Router struct {
}

func RouterSetup(app *fiber.App, h handlers.Handler) {
	api := app.Group("/api")
	{
		api.Get("/get-clients", h.GetAllClients)
		api.Get("/get-crequests", h.GetAllCrequests)
	}

}