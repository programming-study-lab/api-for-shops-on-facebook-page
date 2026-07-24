package photoapp

import (
	"api-for-shops-on-facebook-page/internal/common/config"
	adapter "api-for-shops-on-facebook-page/internal/module/photo_facebook/adapter/photo_facebook_adapter"
	"api-for-shops-on-facebook-page/internal/module/photo_facebook/delivery/http"
	"api-for-shops-on-facebook-page/internal/module/photo_facebook/usecase"
)

type PhotoFacebookApp struct {
	fbConfig *config.FacebookAuthApi
}

func NewPhotoFacebookApp(fbConfig *config.FacebookAuthApi) *PhotoFacebookApp {
	return &PhotoFacebookApp{
		fbConfig: fbConfig,
	}
}

func (config *PhotoFacebookApp) Run(enable string) *http.PhotoFacebookHttp {
	if enable != "enable" {
		return nil
	}

	photoAdapter := adapter.NewPhotoFacebookAdapter(config.fbConfig)
	photoUsecase := usecase.NewPhotoFacebookUsecase(photoAdapter)
	photoHttp := http.NewPhotoFacebookHttp(photoUsecase)

	return photoHttp

}
