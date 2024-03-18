package main

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/handlers"
	"crud-crm/pkg/repository"
	"crud-crm/pkg/routers"

	"github.com/gofiber/fiber/v3"
	log "github.com/sirupsen/logrus"
)

func main() {
	db := database.GetDB()

	mainRepo := repository.NewMainRepository(db)
	handler := handlers.NewHandler(*mainRepo)

	app := fiber.New(fiber.Config{
		ServerHeader: "Storage-CRM",
	})

	router := routers.NewRouter(*handler)
	router.RouterSetup(app)

	log.Fatal(app.Listen(":4000"))
}
