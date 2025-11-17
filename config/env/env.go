package env

import (
	"errors"
	"github.com/HACK3R911/go-todo-api/config"
	"os"
)

const (
	serverPortEnvName = "SRV_PORT"
	dbHostEnvName     = "DB_HOST"
	dbPortEnvName     = "DB_PORT"
	dbUserNameEnvName = "DB_USERNAME"
	dbPasswordEnvName = "DB_PASSWORD"
	dbNameEnvName     = "DB_NAME"
	dbSSLModeEnvName  = "DB_SSLMODE"
)

var _ config.ServerConfig = (*serverConfig)(nil)
var _ config.PostgresConfig = (*postgresConfig)(nil)

type Config struct {
	serverConfig
	postgresConfig
}

type serverConfig struct {
	port string
}

func NewServerConfig() (*serverConfig, error) {
	port := os.Getenv(serverPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("SRV_PORT is not set")
	}

	return &serverConfig{
		port: port,
	}, nil
}

type postgresConfig struct {
	host     string
	port     string
	username string
	password string
	name     string
	sslMode  string
}

func NewPostgresConfig() (*postgresConfig, error) {
	host := os.Getenv(dbHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("DB_HOST is not set")
	}

	port := os.Getenv(dbPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("DB_PORT is not set")
	}

	username := os.Getenv(dbUserNameEnvName)
	if len(username) == 0 {
		return nil, errors.New("DB_USERNAME is not set")
	}

	password := os.Getenv(dbPasswordEnvName)
	if len(password) == 0 {
		return nil, errors.New("DB_PASSWORD is not set")
	}

	name := os.Getenv(dbNameEnvName)
	if len(name) == 0 {
		return nil, errors.New("DB_NAME is not set")
	}

	sslMode := os.Getenv(dbSSLModeEnvName)
	if len(sslMode) == 0 {
		return nil, errors.New("DB_SSLMODE is not set")
	}

	return &postgresConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
		name:     name,
		sslMode:  sslMode,
	}, nil
}

func (cfg *serverConfig) ServerPort() string {
	return cfg.port
}

func (cfg *postgresConfig) DB_Host() string {
	return cfg.host
}

func (cfg *postgresConfig) DB_Port() string {
	return cfg.port
}

func (cfg *postgresConfig) DB_Username() string {
	return cfg.username
}

func (cfg *postgresConfig) DB_Password() string {
	return cfg.password
}

func (cfg *postgresConfig) DB_Name() string {
	return cfg.name
}

func (cfg *postgresConfig) DB_SSLMode() string {
	return cfg.sslMode
}
