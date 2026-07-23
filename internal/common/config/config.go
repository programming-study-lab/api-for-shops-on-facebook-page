package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppUrl struct {
	AppUrl        string `json:"appUrl"`
	AppAPIVersion string `json:"appAPIVersion"`
	AppAddress    string `json:"appAddress"`
	AppPort       string `json:"appPort"`
}

func LoadAppConfig() *AppUrl {
	err := godotenv.Load()

	if err != nil {
		messageError := fmt.Errorf("config.go(LoadAppConfig): %w", err)
		log.Fatalln(messageError)
		return nil
	}

	return &AppUrl{
		AppUrl:        os.Getenv("APP_URL"),
		AppAPIVersion: os.Getenv("APP_API_VERSION"),
		AppAddress:    os.Getenv("APP_ADDRESS"),
		AppPort:       os.Getenv("APP_PORT"),
	}
}
