package main

import (
	"crud-crm/pkg/controllers"
	"crud-crm/pkg/database"
	"crud-crm/pkg/handlers"
	"crud-crm/pkg/repository"
	"crud-crm/pkg/routers"
	"crud-crm/pkg/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	db := database.GetDB()

	mainRepo := repository.NewMainRepository(db)
	controller := controllers.NewControllers(db)
	service := service.NewCrequestSerivce(mainRepo.Crequest)
	handler := handlers.NewHandler(*mainRepo, *controller, *service)

	app := fiber.New(fiber.Config{
		ServerHeader: "Storage-CRM",
	})

	app.Use(cors.New())

	router := routers.NewRouter(*handler)
	router.RouterSetup(app)

	log.Fatal(app.Listen(":4000"))
}
