package main

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/handlers"
	"crud-crm/pkg/repository"
	"crud-crm/pkg/routers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	db := database.GetDB()

	mainRepo := repository.NewMainRepository(db)
	handler := handlers.NewHandler(*mainRepo)

	app := fiber.New(fiber.Config{
		ServerHeader: "Storage-CRM",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router := routers.NewRouter(*handler)
	router.RouterSetup(app)

	log.Fatal(app.Listen(":4000"))
}
