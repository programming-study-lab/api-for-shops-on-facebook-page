package domain

import (
	"context"
)

type PhotoFacebook struct {
	Caption   string
	Url       string
	Published string
}

type PhotoFacebookResponse struct {
	Id *string
}

type PhotoFacebookAdapter interface {
	PhotoCreate(ctx context.Context, photoModel *PhotoFacebook) (*[]byte, error)
}

type PhotoFacebookUsecase interface {
	PhotoCreate(ctx context.Context, photoModel *PhotoFacebook) (interface{}, error)
}
