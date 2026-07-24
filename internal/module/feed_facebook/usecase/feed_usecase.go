package usecase

import (
	"api-for-shops-on-facebook-page/internal/module/feed_facebook/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type feedFacebookUsecase struct {
	adapter domain.FeedFacebookAdapter
}

func NewFeedFacebookUsecase(adapter domain.FeedFacebookAdapter) domain.FeedFacebookUsecase {
	return &feedFacebookUsecase{
		adapter: adapter,
	}
}

func (feed *feedFacebookUsecase) FeedCreate(ctx context.Context, feedFacebook *domain.FeedFacebook) (interface{}, error) {
	// adapter.FeedCreate(ctx, feedFacebook)
	res, err := feed.adapter.FeedCreate(ctx, feedFacebook)
	// return nil, nil

	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_usecase.go(FeedCreate): %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	var result map[string]interface{}

	if err := json.Unmarshal(*res, &result); err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_usecase.go(FeedCreate) Json Decode ล้มเหลว:  %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	return &result, nil
}

func (feed *feedFacebookUsecase) FeedList(ctx context.Context) (interface{}, error) {
	res, err := feed.adapter.FeedList(ctx)
	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_usecase.go(FeedList): %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	var result map[string]interface{}

	if err := json.Unmarshal(*res, &result); err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_usecase.go(FeedList) Json Decode ล้มเหลว:  %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	return &result, nil
}

func (feed *feedFacebookUsecase) FeedUpdate(ctx context.Context, feedId *string, feedFacebook *domain.FeedFacebook) (interface{}, error) {
	// adapter.FeedCreate(ctx, feedFacebook)
	res, err := feed.adapter.FeedUpdate(ctx, feedId, feedFacebook)

	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_usecase.go(FeedUpdate): %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	var result map[string]interface{}

	if err := json.Unmarshal(*res, &result); err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_facebook_adapter.go(FeedUpdate) Json Decode ล้มเหลว:  %w", err)
		log.Fatalln(errorMessage)

		return nil, errorMessage
	}

	return &result, nil
}

// func (feed *feedFacebookUsecase) FeedPhotoCreate(ctx context.Context, feedFacebook *domain.FeedFacebook) (interface{}, error) {
// 	// adapter.FeedCreate(ctx, feedFacebook)
// 	res, err := feed.adapter.FeedCreate(ctx, feedFacebook)

// 	if err != nil {
// 		errorMessage := fmt.Errorf("[WARNING] feed_usecase.go(FeedCreate): %w", err)
// 		log.Fatalln(errorMessage)

// 		return nil, errorMessage
// 	}

// 	var result map[string]interface{}

// 	if err := json.Unmarshal(*res, &result); err != nil {
// 		errorMessage := fmt.Errorf("[WARNING] feed_facebook_adapter.go(FeedCreate) Json Decode ล้มเหลว:  %w", err)
// 		log.Fatalln(errorMessage)

// 		return nil, errorMessage
// 	}

// 	return &result, nil
// }
