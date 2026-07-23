package commentdto

import "api-for-shops-on-facebook-page/internal/module/comment_facebook/domain"

type CommentFacebookDTO struct {
	// FeedId string `json:"feedId" form:"feed_id" binding:"required"`
	FeedId string `json:"feedId" form:"feed_id"`
}

func (commentModelDTO *CommentFacebookDTO) ToDomain() *domain.CommentFacebook {
	return &domain.CommentFacebook{
		FeedId: commentModelDTO.FeedId,
	}
}

func (*CommentFacebookDTO) FromDomain(commentModel *domain.CommentFacebook) *CommentFacebookDTO {
	return &CommentFacebookDTO{
		FeedId: commentModel.FeedId,
	}
}
