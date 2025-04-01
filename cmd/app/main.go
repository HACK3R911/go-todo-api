package main

import (
	"context"
	"github.com/HACK3R911/go-todo-api"
	"github.com/HACK3R911/go-todo-api/pkg/handler"
	"github.com/HACK3R911/go-todo-api/pkg/repository"
	"github.com/HACK3R911/go-todo-api/pkg/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error: initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error: loading .env: %s", err.Error())
	}

	db, err := repository.NewPostgreDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),

		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error: initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
