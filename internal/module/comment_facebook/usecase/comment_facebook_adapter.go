package usecase

import (
	"api-for-shops-on-facebook-page/internal/module/comment_facebook/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type commentFacebookUsecase struct {
	adapter domain.CommentFacebookAdapter
}

func NewCommentFacebookUsecase(adapter domain.CommentFacebookAdapter) domain.CommentFacebookUsecase {
	return &commentFacebookUsecase{
		adapter: adapter,
	}
}

func (apt *commentFacebookUsecase) CommentList(ctx context.Context, commentModel *domain.CommentFacebook) (interface{}, error) {
	res, err := apt.adapter.CommentList(ctx, commentModel)
	if err != nil {
		messageErrro := fmt.Errorf("coment_facebook_usecase.go(CommentList): %w", err)
		log.Fatalln(messageErrro)
		return nil, messageErrro
	}

	var result map[string]interface{}

	if err := json.Unmarshal(*res, &result); err != nil {
		messageError := fmt.Errorf("comment_facebook_adapter.go(CommentList): %w", err)
		log.Fatalln(messageError)
		return nil, messageError
	}

	return &result, nil

}
