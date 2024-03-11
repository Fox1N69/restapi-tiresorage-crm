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

	router := routers.NewRouter(handlers.Handler{})
	router.RouterSetup(app)

	log.Fatal(app.Listen(":4000"))
}

