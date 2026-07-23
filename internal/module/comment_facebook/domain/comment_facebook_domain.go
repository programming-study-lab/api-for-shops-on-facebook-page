package domain

import "context"

type CommentFacebook struct {
	FeedId    string
	CommentId string
}

type CommentFacebookAdapter interface {
	CommentList(ctx context.Context, commentModel *CommentFacebook) (*[]byte, error)
}

type CommentFacebookUsecase interface {
	CommentList(ctx context.Context, commentModel *CommentFacebook) (interface{}, error)
}
