package main

import (
	"github.com/HACK3R911/go-todo-api"
	"github.com/HACK3R911/go-todo-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
