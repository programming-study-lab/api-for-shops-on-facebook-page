package facebookservices

import (
	"api-for-shops-on-facebook-page/configs"
	"fmt"

	"github.com/huandu/facebook/v2"
)

func FacebookInit() (*facebook.Session, error) {
	fbConfig := configs.FacebookConfig()
	globalApp := facebook.New(fbConfig.FacebookPageId, fbConfig.AccessToken)

	session := globalApp.Session(fbConfig.AccessToken)

	err := session.Validate()

	if err == nil {
		return session, nil
	} else {
		return nil, fmt.Errorf("%s", err.Error())
	}

}
