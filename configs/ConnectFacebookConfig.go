package configs

import (
	"os"

	fb "github.com/huandu/facebook/v2"
	"github.com/joho/godotenv"
)

type FacebookAuthApi struct {
	VersionGraph   string `json:"versionGraph"`
	AccessToken    string `json:"AccessToken"`
	FacebookPageId string `json:"facebookPageId"`
	FacebookAPi    string `json:"facebookApi"`
}

func ConnectFacebookGraphApi() *FacebookAuthApi {
	godotenv.Load()

	fb.Version = os.Getenv("FACEBOOK_GRAPH_VERSION")

	return &FacebookAuthApi{
		VersionGraph:   os.Getenv("FACEBOOK_GRAPH_VERSION"),
		AccessToken:    os.Getenv("FACEBOOK_ACCESS_TOKEN"),
		FacebookPageId: os.Getenv("FACEBOOK_PAGE_ID"),
		FacebookAPi:    os.Getenv("FACEBOOK_API"),
	}
}
