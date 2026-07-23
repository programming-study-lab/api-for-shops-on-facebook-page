package commentfacebookadapter

import (
	"api-for-shops-on-facebook-page/internal/common/config"
	"api-for-shops-on-facebook-page/internal/module/comment_facebook/domain"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CommentFacebookAdapter struct {
	fbConfig *config.FacebookAuthApi
}

func NewCommentFacebookAdapter(fbConfig *config.FacebookAuthApi) domain.CommentFacebookAdapter {
	return &CommentFacebookAdapter{
		fbConfig: fbConfig,
	}
}

func (config *CommentFacebookAdapter) CommentList(ctx context.Context, commentModel *domain.CommentFacebook) (*[]byte, error) {

	var params string

	if commentModel.FeedId != "" {
		params = commentModel.FeedId
	} else if commentModel.CommentId != "" {
		params = commentModel.CommentId
	}

	url := config.fbConfig.FacebookAPI
	url += "/" + config.fbConfig.VersionGraph
	url += "/" + params
	url += "/" + "comments"
	url += "?access_token=" + config.fbConfig.AccessToken

	// res, err := http.Get(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errorMessage := fmt.Errorf("comment_facebook_adapter.go(CommentList) %w", err)
		log.Fatalln(errorMessage)
		return nil, errorMessage
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		errorMessage := fmt.Errorf("comment_facebook_adapter.go(CommentList) %w", err)
		log.Fatalln(errorMessage)
		return nil, errorMessage
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	return &body, nil
}
