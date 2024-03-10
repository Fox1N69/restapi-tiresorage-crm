package main

import (
	"context"

	server "github.com/Fox1N69/rest-tsc"
	handler "github.com/Fox1N69/rest-tsc/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(server.Server)

	log.Fatal(srv.Run("8000", handlers.InitRouting()))
	log.Fatal(srv.Shutdown(context.Background()))
}
