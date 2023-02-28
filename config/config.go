package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		errors.New("Error loading .env file")
	}
	config := Config{Port: os.Getenv("APP_PORT")}
	return config, err
}
