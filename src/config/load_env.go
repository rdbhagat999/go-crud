package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST         string        `mapstructure:"DB_HOST"`
	DB_USER         string        `mapstructure:"DB_USER"`
	DB_PASSWORD     string        `mapstructure:"DB_PASSWORD"`
	DB_NAME         string        `mapstructure:"DB_NAME"`
	DB_PORT         string        `mapstructure:"DB_PORT"`
	SERVER_PORT     string        `mapstructure:"SERVER_PORT"`
	TOKEN_EXPIRE_IN time.Duration `mapstructure:"TOKEN_EXPIRE_IN"`
	TOKEN_MAX_AGE   string        `mapstructure:"TOKEN_MAX_AGE"`
	TOKEN_SECRET    string        `mapstructure:"TOKEN_SECRET"`
	GITHUB_USER     string        `mapstructure:"GITHUB_USER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
