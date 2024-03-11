package main

import (
	"crud-crm/pkg/database"
	"crud-crm/pkg/handlers"
	"crud-crm/pkg/routers"

	"github.com/gofiber/fiber/v3"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
	})

	log.Infoln("database connect...", database.InitDB())
	routers.RouterSetup(app, handlers.Handler{})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	log.Fatal(app.Listen(":4000"))
}
