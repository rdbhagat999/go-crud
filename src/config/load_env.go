package config

import (
	"time"

	"github.com/joho/godotenv"
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

func LoadConfig(path string) (cfg Config, err error) {
	loadErr := godotenv.Load()
	var config Config

	if loadErr != nil {
		return config, loadErr
	}

	var myEnv map[string]string
	myEnv, err = godotenv.Read()

	if err != nil {
		return config, err
	}

	DB_HOST := myEnv["DB_HOST"]
	DB_USER := myEnv["DB_USER"]
	DB_PASSWORD := myEnv["DB_PASSWORD"]
	DB_NAME := myEnv["DB_NAME"]
	DB_PORT := myEnv["DB_PORT"]
	SERVER_PORT := myEnv["SERVER_PORT"]
	TOKEN_MAX_AGE := myEnv["TOKEN_MAX_AGE"]
	TOKEN_SECRET := myEnv["TOKEN_SECRET"]
	GITHUB_USER := myEnv["GITHUB_USER"]

	config.DB_HOST = DB_HOST
	config.DB_USER = DB_USER
	config.DB_PASSWORD = DB_PASSWORD
	config.DB_NAME = DB_NAME
	config.DB_PORT = DB_PORT
	config.SERVER_PORT = SERVER_PORT
	config.TOKEN_MAX_AGE = TOKEN_MAX_AGE
	config.TOKEN_SECRET = TOKEN_SECRET
	config.GITHUB_USER = GITHUB_USER

	return config, err
}
