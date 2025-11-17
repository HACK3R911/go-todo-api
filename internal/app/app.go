package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HACK3R911/go-todo-api/config"
	"github.com/HACK3R911/go-todo-api/config/env"
	"github.com/HACK3R911/go-todo-api/internal/handler"
	"github.com/HACK3R911/go-todo-api/internal/repository"
	"github.com/HACK3R911/go-todo-api/internal/service"
	"github.com/HACK3R911/go-todo-api/pkg/server"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger *logrus.Logger
	server *server.Server
	db     *sqlx.DB
	cfg    *config.Config
}

func New(ctx context.Context, configPath string) (*App, error) {
	app := &App{}

	err := config.Load(configPath)
	if err != nil {
		return nil, err
	}

	serverConfig, err := env.NewServerConfig()
	if err != nil {
		return nil, err
	}

	postgresConfig, err := env.NewPostgresConfig()
	if err != nil {
		return nil, err
	}

	app.cfg = &config.Config{
		ServerConfig:   serverConfig,
		PostgresConfig: postgresConfig,
	}

	//log, err := logger.New(app.cfg.ServerConfig.ServerPort())
	//if err != nil {
	//	return nil, err
	//}
	//app.logger = log

	db, err := repository.NewPostgreDB(repository.Config{
		Host:     postgresConfig.DBHost(),
		Port:     postgresConfig.DBPort(),
		Username: postgresConfig.DBUsername(),
		Password: postgresConfig.DBPassword(),
		DBName:   postgresConfig.DBName(),
		SSLMode:  postgresConfig.DBSSLMode(),
	})
	if err != nil {
		return nil, err
	}
	app.db = db

	return app, nil
}

func (a *App) Run() error {
	defer a.db.Close()

	repo := repository.NewRepository(a.db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	a.server = new(server.Server)

	go a.startServer(handler)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	sig := <-quit
	logrus.Infof("signal %s received", sig.String())

	shutdownServer, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := a.shutdownServer(shutdownServer)
	if err != nil {
		logrus.Errorf("error shutdown server: %s", err.Error())
	}

	return nil
}

func (a *App) startServer(handler *handler.Handler) error {
	logrus.Infof("starting server on port %s...", a.cfg.ServerConfig.ServerPort())

	err := a.server.Run(a.cfg.ServerConfig.ServerPort(), handler.InitRoutes())
	if err != nil {
		logrus.Fatalf("error starting server: %s", err.Error())
	} else {
		logrus.Infof("server started on port %s", a.cfg.ServerConfig.ServerPort())
	}

	return nil
}

func (a *App) shutdownServer(ctx context.Context) error {
	logrus.Info("shutdown server...")
	if err := a.server.Shutdown(ctx); err != nil {
		logrus.Errorf("error shutdown server: %s", err.Error())
	} else {
		logrus.Info("server stopped.")
	}

	if err := a.db.Close(); err != nil {
		logrus.Errorf("error on db connection close: %s", err.Error())
	} else {
		logrus.Info("db closed.")
	}

	return nil
}
