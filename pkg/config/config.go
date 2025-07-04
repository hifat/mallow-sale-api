package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type App struct {
	Name string `mapstructure:"APP_NAME"`
	Host string `mapstructure:"APP_HOST"`
	Port string `mapstructure:"APP_PORT"`
	Mode string `mapstructure:"APP_MODE"`
}

type DB struct {
	DBName   string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	SSLMode  string `mapstructure:"DB_SSL_MODE"`
	Schema   string `mapstructure:"DB_SCHEMA"`
}

type Config struct {
	App App
	DB  DB
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		path = ".env"
	}

	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		// .env file is optional, so only log if not found
		log.Printf("No .env file found or error reading it: %v", err)
	}

	var cfg Config
	errCh := make(chan error, 2)

	go func() {
		if err := viper.Unmarshal(&cfg.App); err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	go func() {
		if err := viper.Unmarshal(&cfg.DB); err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
