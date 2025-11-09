package config

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type App struct {
	Name string
	Host string
	Port string
	Mode string
}

type DB struct {
	Protocol string
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

type Auth struct {
	AccessTokenSecret   string
	AccessTokenExpires  time.Duration
	RefreshTokenSecret  string
	RefreshTokenExpires time.Duration
}

type Config struct {
	App  App
	DB   DB
	GRPC GRPC
	Auth Auth
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
			Protocol: viper.GetString("DB_PROTOCOL"),
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
		Auth: Auth{
			AccessTokenSecret:   viper.GetString("ACCESS_TOKEN_SECRET"),
			AccessTokenExpires:  viper.GetDuration("ACCESS_TOKEN_EXPIRES"),
			RefreshTokenSecret:  viper.GetString("REFRESH_TOKEN_SECRET"),
			RefreshTokenExpires: viper.GetDuration("REFRESH_TOKEN_EXPIRES"),
		},
	}

	return &cfg, nil
}
