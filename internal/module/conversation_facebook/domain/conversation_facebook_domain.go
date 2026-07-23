package domain

import "context"

type ConversationFacebook struct {
	FeedId    string
	CommentId string
}

type ConversationFacebookAdapter interface {
	CommentList(ctx context.Context, conversationModel *ConversationFacebook) (*[]byte, error)
}

type ConversationFacebookUsecase interface {
	CommentList(ctx context.Context, conversationModel *ConversationFacebook) (interface{}, error)
}
