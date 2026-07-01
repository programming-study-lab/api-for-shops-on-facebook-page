package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type FacebookAuthApi struct {
	VersionGraph         string `json:"versionGraph"`
	AccessToken          string `json:"AccessToken"`
	FacebookPageId       string `json:"facebookPageId"`
	FacebookAPI          string `json:"facebookApi"`
	FacebookWebhookToken string `json:"facebookWebhookToken"`
}

func FacebookConfig() *FacebookAuthApi {
	err := godotenv.Load()

	if err != nil {
		log.Println("ไม่พบไฟล์ .env")
	}

	return &FacebookAuthApi{
		VersionGraph:         os.Getenv("FACEBOOK_GRAPH_VERSION"),
		AccessToken:          os.Getenv("FACEBOOK_ACCESS_TOKEN"),
		FacebookPageId:       os.Getenv("FACEBOOK_PAGE_ID"),
		FacebookAPI:          os.Getenv("FACEBOOK_API"),
		FacebookWebhookToken: os.Getenv("FACEBOOK_WEBHOOK_TOKEN"),
	}
}
