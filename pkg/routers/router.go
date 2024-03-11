package routers

import (
	"crud-crm/pkg/handlers"
	"crud-crm/pkg/repository"

	"github.com/gofiber/fiber/v3"
)

// создание структуры router с полем handler
type Router struct {
	handler handlers.Handler
	repo repository.ClientRepository
}

//новый экземпляр типа router с обработчиком handler
func NewRouter(h handlers.Handler) *Router {
	return &Router{handler: h}
}

//метод запускается на экзмепляре router и устанавливает маршруты для fiber.app
func (r *Router) RouterSetup(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Get("/clients", r.handler.GetAllClients)
		api.Get("/get-crequests", r.handler.GetAllCrequests)
	}

}
