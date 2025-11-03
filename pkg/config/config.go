package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type App struct {
	Name string
	Host string
	Port string
	Mode string
}

type DB struct {
	DBName   string
	Username string
	Password string
	Host     string
	Port     string
	SSLMode  string
	Schema   string
}

type GRPC struct {
	Host string
	Port string
}

type Config struct {
	App  App
	DB   DB
	GRPC GRPC
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		path = ".env"
	}

	viper.SetConfigFile(path)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No .env file found, using environment variables")
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfg := Config{
		App: App{
			Name: viper.GetString("APP_NAME"),
			Host: viper.GetString("APP_HOST"),
			Port: viper.GetString("APP_PORT"),
			Mode: viper.GetString("APP_MODE"),
		},
		DB: DB{
			DBName:   viper.GetString("DB_NAME"),
			Username: viper.GetString("DB_USERNAME"),
			Password: viper.GetString("DB_PASSWORD"),
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			SSLMode:  viper.GetString("DB_SSL_MODE"),
			Schema:   viper.GetString("DB_SCHEMA"),
		},
		GRPC: GRPC{
			Host: viper.GetString("GRPC_HOST"),
			Port: viper.GetString("GRPC_PORT"),
		},
	}

	return &cfg, nil
}
