package photofacebookadapter

import (
	"api-for-shops-on-facebook-page/internal/common/config"
	photofacebookdto "api-for-shops-on-facebook-page/internal/module/photo_facebook/adapter/data_transfer_object"
	"api-for-shops-on-facebook-page/internal/module/photo_facebook/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PhotoFacebookAdapter struct {
	fbConfig *config.FacebookAuthApi
}

func NewPhotoFacebookAdapter(fbConfig *config.FacebookAuthApi) domain.PhotoFacebookAdapter {
	return &PhotoFacebookAdapter{
		fbConfig: fbConfig,
	}
}

func (fbConfig *PhotoFacebookAdapter) PhotoCreate(ctx context.Context, photoModel *domain.PhotoFacebook) (*[]byte, error) {
	url := fbConfig.fbConfig.FacebookAPI + "/" + fbConfig.fbConfig.VersionGraph
	url += "/" + "me"
	url += "/" + "photos?"
	url += "access_token=" + fbConfig.fbConfig.AccessToken

	photoDTO := &photofacebookdto.PhotoFacebookDTO{}

	photo := photoDTO.FromDomain(photoModel)

	photoBodyData := new(bytes.Buffer)

	err := json.NewEncoder(photoBodyData).Encode(&photo)
	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] photo_facebook_adapter.go(PhotoCreate) เข้ารหัส Encoder ล้มเหลว: %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	res, err := http.Post(url, "application/json", photoBodyData)
	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] photo_facebook_adapter.go(PhotoCreate) request ไปยัง Facebook ล้มเหลว:  %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	return &body, nil
}
