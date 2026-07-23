package domain

import (
	"context"
)

type FeedFacebook struct {
	Message   string
	Published bool
	Privacy   *PrivacyParams
}

type PrivacyParams struct {
	Value string
}

type FeedFacebookUsecase interface {
	FeedCreate(ctx context.Context, feedFacebook *FeedFacebook) (interface{}, error)
	FeedList(ctx context.Context) (interface{}, error)
	FeedUpdate(ctx context.Context, feedId *string, feedFacebook *FeedFacebook) (interface{}, error)
}

type FeedFacebookAdapter interface {
	FeedCreate(ctx context.Context, feedBody *FeedFacebook) (*[]byte, error)
	FeedList(ctx context.Context) (*[]byte, error)
	FeedUpdate(ctx context.Context, feedId *string, feedBody *FeedFacebook) (*[]byte, error)
}
