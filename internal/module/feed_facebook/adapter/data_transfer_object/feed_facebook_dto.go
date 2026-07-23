package feedfacebookdto

import "api-for-shops-on-facebook-page/internal/module/feed_facebook/domain"

type FeedFacebookDTO struct {
	Message   string            `json:"message"`
	Published bool              `json:"published"`
	Privacy   *PrivacyParamsDTO `json:"privacy"`
}

type PrivacyParamsDTO struct {
	Value string `json:"value"`
}

func (feedDTO *FeedFacebookDTO) ToDomain() *domain.FeedFacebook {
	return &domain.FeedFacebook{
		Message:   feedDTO.Message,
		Published: feedDTO.Published,
		Privacy: &domain.PrivacyParams{
			Value: feedDTO.Privacy.Value,
		},
	}
}

func (feedDTO *FeedFacebookDTO) FromDomain(feedDomain *domain.FeedFacebook) *FeedFacebookDTO {

	return &FeedFacebookDTO{
		Message:   feedDomain.Message,
		Published: feedDomain.Published,
		Privacy: &PrivacyParamsDTO{
			Value: feedDomain.Privacy.Value,
		},
	}

}
