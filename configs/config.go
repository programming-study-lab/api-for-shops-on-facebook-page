package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Domain struct {
	URL_Host  string `json:"url_host"`
	Port_Host string `json:"port_host"`
}

func LoadEnv() *Domain {
	godotenv.Load()

	return &Domain{
		URL_Host:  os.Getenv("URL_HOST"),
		Port_Host: os.Getenv("PORT_HOST"),
	}
}
