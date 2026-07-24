package feedfacebookdto

import (
	"api-for-shops-on-facebook-page/internal/module/feed_facebook/domain"
)

type FeedFacebookDTO struct {
	Message               *string           `json:"message" form:"message"`
	Published             *bool             `json:"published" form:"published"`
	Privacy               *PrivacyParamsDTO `json:"privacy" form:"privacy"`
	AttachedMediaAttached *string           `json:"attached_media" form:"attached_media"`
}

type PrivacyParamsDTO struct {
	Value *string `json:"value"`
}

type AttachedMediaAttachedParams struct {
	MediaFBId *string `json:"media_fbid" form:"media_fbid"`
}

func (feedDTO *FeedFacebookDTO) ToDomain() *domain.FeedFacebook {
	feedDomain := &domain.FeedFacebook{}

	if feedDTO.Message != nil {
		feedDomain.Message = feedDTO.Message
	}

	if feedDTO.Published != nil {
		feedDomain.Published = feedDTO.Published
	}

	if feedDTO.Privacy != nil {

		feedDomain.Privacy = &domain.PrivacyParams{
			Value: feedDTO.Privacy.Value,
		}

	}

	if feedDTO.AttachedMediaAttached != nil {
		feedDomain.AttachedMediaAttached = feedDTO.AttachedMediaAttached
	}

	return feedDomain

}

func (feedDTO *FeedFacebookDTO) FromDomain(feedDomain *domain.FeedFacebook) *FeedFacebookDTO {
	if feedDomain.Message != nil {
		feedDTO.Message = feedDomain.Message
	}
	if feedDomain.Published != nil {
		feedDTO.Published = feedDomain.Published
	}

	if feedDomain.Privacy != nil {
		feedDTO.Privacy = &PrivacyParamsDTO{
			Value: feedDomain.Privacy.Value,
		}
	}

	if feedDomain.AttachedMediaAttached != nil {
		feedDTO.AttachedMediaAttached = feedDomain.AttachedMediaAttached
	}

	return feedDTO

}
