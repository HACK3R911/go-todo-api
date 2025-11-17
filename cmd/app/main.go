package main

import (
	"context"
	"flag"
	"github.com/HACK3R911/go-todo-api"
	"github.com/HACK3R911/go-todo-api/config"
	"github.com/HACK3R911/go-todo-api/config/env"
	"github.com/HACK3R911/go-todo-api/internal/handler"
	"github.com/HACK3R911/go-todo-api/internal/repository"
	"github.com/HACK3R911/go-todo-api/internal/service"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "путь к конфигурационному файлу")
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	flag.Parse()

	if err := config.Load(configPath); err != nil {
		logrus.Fatalf("error: loading .env: %s", err.Error())
	}

	serverConfig, err := env.NewServerConfig()
	if err != nil {
		logrus.Fatalf("error: loading server config: %s", err.Error())
	}

	postgresConfig, err := env.NewPostgresConfig()
	if err != nil {
		logrus.Fatalf("error: loading postgres config: %s", err.Error())
	}

	db, err := repository.NewPostgreDB(repository.Config{
		Host:     postgresConfig.DB_Host(),
		Port:     postgresConfig.DB_Port(),
		Username: postgresConfig.DB_Username(),
		Password: postgresConfig.DB_Password(),
		DBName:   postgresConfig.DB_Name(),
		SSLMode:  postgresConfig.DB_SSLMode(),
	})
	if err != nil {
		logrus.Fatalf("error: initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(serverConfig.ServerPort(), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error running server: %s", err.Error())
		}
	}()

	logrus.Println("TodoApp is started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("TodoApp is stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error on server shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error on db connection close: %s", err.Error())
	}
}
