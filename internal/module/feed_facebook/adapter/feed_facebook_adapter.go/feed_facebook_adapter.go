// package feedfacebookdelivergo
package adapter

import (
	"api-for-shops-on-facebook-page/internal/common/config"
	feedfacebookdto "api-for-shops-on-facebook-page/internal/module/feed_facebook/adapter/data_transfer_object"
	"api-for-shops-on-facebook-page/internal/module/feed_facebook/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type FeedFacebookAdapter struct {
	fbConfig *config.FacebookAuthApi
}

func NewFeedFacebookAdapter(facebookConfig *config.FacebookAuthApi) domain.FeedFacebookAdapter {
	return &FeedFacebookAdapter{
		fbConfig: facebookConfig,
	}
}

func (fbConfig *FeedFacebookAdapter) FeedCreate(ctx context.Context, feedBody *domain.FeedFacebook) (*[]byte, error) {
	url := fbConfig.fbConfig.FacebookAPI + "/" + fbConfig.fbConfig.VersionGraph
	url += "/" + "me"
	url += "/" + "feed?"
	url += "access_token=" + fbConfig.fbConfig.AccessToken
	// url += "&" + ""

	feedDTO := &feedfacebookdto.FeedFacebookDTO{}

	feed := feedDTO.FromDomain(feedBody)

	// fmt.Printf("\n[debug] %s\n", feed)

	feedBodyData := new(bytes.Buffer)

	err := json.NewEncoder(feedBodyData).Encode(&feed)
	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_facebook_adapter.go(FeedCreate) เข้ารหัส Encoder ล้มเหลว: %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	res, err := http.Post(url, "application/json", feedBodyData)
	if err != nil {
		// fmt.Printf("\nfeed_facebook_adapter.go(FeedCreate) request ไปยัง Facebook ล้มเหลว: %s", err.Error())
		errorMessage := fmt.Errorf("[WARNING] feed_facebook_adapter.go(FeedCreate) request ไปยัง Facebook ล้มเหลว:  %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	// var result map[string]interface{}

	// json.NewDecoder(res.Body).Decode(&body)

	return &body, nil
}

func (config *FeedFacebookAdapter) FeedList(ctx context.Context) (*[]byte, error) {
	url := config.fbConfig.FacebookAPI + "/" + config.fbConfig.VersionGraph
	url += "/" + "me"
	url += "/" + "feed"
	url += "?" + "fields=" + "id,message,created_time,full_picture,attachments"
	url += "&" + "access_token=" + config.fbConfig.AccessToken

	res, err := http.Get(url)
	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_facebook_adapter.go(FeedList) request ไปยัง Facebook ล้มเหลว:  %w", err)
		log.Fatalln(errorMessage)

		if res.StatusCode < 200 || res.StatusCode >= 300 {
			fmt.Printf("Server returned error status: %d\n", res.StatusCode)
			// return nil, nil
		}

		return nil, errorMessage
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	return &body, nil
}

func (fbConfig *FeedFacebookAdapter) FeedUpdate(ctx context.Context, feedId *string, feedBody *domain.FeedFacebook) (*[]byte, error) {

	url := fbConfig.fbConfig.FacebookAPI
	url += "/" + *feedId
	url += "?" + "access_token=" + fbConfig.fbConfig.AccessToken

	feedDTO := &feedfacebookdto.FeedFacebookDTO{}
	feed := feedDTO.FromDomain(feedBody)

	feedBodyData := new(bytes.Buffer)

	err := json.NewEncoder(feedBodyData).Encode(&feed)
	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_facebook_adapter.go(FeedCreate) เข้ารหัส Encoder ล้มเหลว: %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	res, err := http.Post(url, "application/json", feedBodyData)
	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_facebook_adapter.go(FeedCreate) request ไปยัง Facebook ล้มเหลว:  %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	return &body, nil
}
