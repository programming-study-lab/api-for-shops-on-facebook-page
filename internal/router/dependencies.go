package router

import (
	comment "api-for-shops-on-facebook-page/internal/module/comment_facebook/delivery/http"
	feed "api-for-shops-on-facebook-page/internal/module/feed_facebook/delivery/http"
)

type Dependencies struct {
	Feed    *feed.FeetFacebookHttp
	Comment *comment.CommentFacebookHttp
}

func NewDependencies() *Dependencies {
	return &Dependencies{}
}
