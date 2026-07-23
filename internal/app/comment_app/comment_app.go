package commentapp

import (
	"api-for-shops-on-facebook-page/internal/common/config"
	adapter "api-for-shops-on-facebook-page/internal/module/comment_facebook/adapter/comment_facebook_adapter"
	"api-for-shops-on-facebook-page/internal/module/comment_facebook/delivery/http"
	"api-for-shops-on-facebook-page/internal/module/comment_facebook/usecase"
)

type CommentFacebookApp struct {
	fbConfig *config.FacebookAuthApi
}

func NewCommentFacebookApp(fbConfig *config.FacebookAuthApi) *CommentFacebookApp {
	return &CommentFacebookApp{
		fbConfig: fbConfig,
	}
}

func (config *CommentFacebookApp) Run(enable string) *http.CommentFacebookHttp {
	if enable != "enable" {
		return nil
	}

	commentAdapter := adapter.NewCommentFacebookAdapter(config.fbConfig)
	commentUsecase := usecase.NewCommentFacebookUsecase(commentAdapter)
	commentHttp := http.NewCommentFacebookHttp(commentUsecase)

	return commentHttp

}
