package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App  App
	Db   Db
	Auth Auth
	GRPC GRPC
}

type App struct {
	Service string `mapstructure:"APP_SERVICE"`
	Host    string `mapstructure:"APP_HOST"`
	Port    string `mapstructure:"APP_PORT"`
	Name    string `mapstructure:"APP_NAME"`
}

type Db struct {
	Host     string `mapstructure:"DB_HOST"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Port     string `mapstructure:"DB_PORT"`
	Name     string `mapstructure:"DB_NAME"`
}

type Auth struct {
	AccessToken         string        `mapstructure:"ACCESS_TOKEN"`
	AccessTokenExpires  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRES"`
	RefreshToken        string        `mapstructure:"REFRESH_TOKEN"`
	RefreshTokenExpires time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRES"`
	APIKey              string        `mapstructure:"API_KEY"`
}
type GRPC struct {
	InventoryHost string `mapstructure:"GRPC_INVENTORY_HOST"`
	UsageUnitHost string `mapstructure:"GRPC_USAGE_UNIT_HOST"`
}

func (c *Config) Init(path string, filename string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.Unmarshal(&c.App)
	viper.Unmarshal(&c.Db)
	viper.Unmarshal(&c.GRPC)

	return nil
}

func LoadAppConfig(paths string) *Config {
	split := strings.Split(paths, "/")
	path := strings.Join(split[:len(split)-1], "/")
	filename := split[len(split)-1]

	cfg := Config{}

	err := cfg.Init(path, filename)
	if err != nil {
		panic(err)
	}

	return &cfg
}
