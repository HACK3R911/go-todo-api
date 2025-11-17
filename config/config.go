package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Load(path string) error {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		logrus.Fatal("Ошибка загрузки .env файла: ", err)
	}

	settings := v.AllSettings()
	for key := range settings {
		value := v.GetString(key)
		os.Setenv(key, value)
	}

	return nil
}

type ServerConfig interface {
	ServerPort() string
}

type PostgresConfig interface {
	DB_Host() string
	DB_Port() string
	DB_Username() string
	DB_Password() string
	DB_Name() string
	DB_SSLMode() string
}
