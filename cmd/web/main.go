package main

import (
	"github.com/HACK3R911/go-todo-api"
	"github.com/HACK3R911/go-todo-api/pkg/handler"
	"github.com/HACK3R911/go-todo-api/pkg/repository"
	"github.com/HACK3R911/go-todo-api/pkg/service"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error: initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
