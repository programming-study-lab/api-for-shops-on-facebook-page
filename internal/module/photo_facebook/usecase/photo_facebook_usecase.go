package usecase

import (
	"api-for-shops-on-facebook-page/internal/module/photo_facebook/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type photoFacebookUsecase struct {
	adapter domain.PhotoFacebookAdapter
}

func NewPhotoFacebookUsecase(adapter domain.PhotoFacebookAdapter) domain.PhotoFacebookUsecase {
	return &photoFacebookUsecase{
		adapter: adapter,
	}
}

func (photo *photoFacebookUsecase) PhotoCreate(ctx context.Context, photoModel *domain.PhotoFacebook) (interface{}, error) {
	// adapter.FeedCreate(ctx, feedFacebook)
	fmt.Printf("\n[photo_facebook_usecase.go(PhotoCreate)] %s\n", photoModel)
	// return nil, nil
	res, err := photo.adapter.PhotoCreate(ctx, photoModel)

	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] photo_facebook_usecase.go(PhotoCreate): %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	var result map[string]interface{}

	if err := json.Unmarshal(*res, &result); err != nil {
		errorMessage := fmt.Errorf("[WARNING] photo_facebook_usecase.go(PhotoCreate) Json Decode ล้มเหลว:  %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	return &result, nil
}
