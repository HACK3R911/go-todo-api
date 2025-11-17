package main

import (
	"context"
	"flag"
	"github.com/HACK3R911/go-todo-api/internal/app"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "путь к конфигурационному файлу")
}

func main() {
	ctx := context.Background()

	logrus.SetFormatter(new(logrus.JSONFormatter))
	flag.Parse()

	application, err := app.New(ctx, configPath)
	if err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}
