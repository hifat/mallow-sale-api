package config

import "github.com/spf13/viper"

type Config struct {
	App App
	Db  Db
}

type App struct {
	Host string `mapstructure:"APP_HOST"`
	Port string `mapstructure:"APP_PORT"`
}

type Db struct {
	Host     string `mapstructure:"DB_HOST"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Port     string `mapstructure:"DB_PORT"`
	Name     string `mapstructure:"DB_NAME"`
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

	return nil
}

func LoadAppConfig(path string, filename string) *Config {
	cfg := Config{}

	err := cfg.Init(path, filename)
	if err != nil {
		panic(err)
	}

	return &cfg
}
