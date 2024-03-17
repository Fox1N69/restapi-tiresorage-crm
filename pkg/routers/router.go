package routers

import (
	"crud-crm/pkg/handlers"
	"crud-crm/pkg/middlewares"

	//"crud-crm/pkg/repository"

	"github.com/gofiber/fiber/v3"
)

// создание структуры router с полем handler
type Router struct {
	handler     handlers.Handler
	middlewares *middlewares.Middlewares
	//mainRepo repository.MainRepository
}

// новый экземпляр типа router с обработчиком handler
func NewRouter(h handlers.Handler) *Router {
	return &Router{handler: h}
}

// метод запускается на экзмепляре router и устанавливает маршруты для fiber.app
func (r *Router) RouterSetup(app *fiber.App) {
	auth := app.Group("/auth")
	{
		auth.Use(r.middlewares.Auth)
		auth.Post("/register", r.handler.Register)
		auth.Post("/login", r.handler.Login)
		auth.Post("/logout", r.handler.Logout)
	}

	api := app.Group("/api")
	{
		client := api.Group("/client")
		{
			client.Get("/", r.handler.GetAllClients)
			client.Post("/set", r.handler.CreateClient)
			client.Get("/:id", r.handler.GetClientByID)
			client.Put("/:id", r.handler.UpdateClient)
			client.Delete("/:id", r.handler.DeleteClient)
		}
		crequest := api.Group("/crequest")
		{
			crequest.Get("/", r.handler.GetAllCrequests)
			crequest.Post("/set", r.handler.CreateCrequest)
			crequest.Get("/:id", r.handler.GetCrequestByID)
			crequest.Put("/:id", r.handler.UpdateCrequest)
			crequest.Delete("/:id", r.handler.DeleteCrequest)
		}

	}
}
