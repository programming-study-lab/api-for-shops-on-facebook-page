package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppUrl struct {
	AppUrl  string `json:"appUrl"`
	AppPort string `json:"appPort"`
}

func LoadEnv() *AppUrl {
	godotenv.Load()

	return &AppUrl{
		AppUrl:  os.Getenv("APP_URL"),
		AppPort: os.Getenv("APP_PORT"),
	}
}
