package v1

import (
	"api-for-shops-on-facebook-page/internal/module/feed_facebook/delivery/http"

	"github.com/gin-gonic/gin"
)

func FeedFacebookRouter(groupV1 *gin.RouterGroup, handler *http.FeetFacebookHttp) {
	feed := groupV1.Group("/feed")
	{
		feed.POST("/", handler.FeedCreate)
		feed.GET("/", handler.FeedList)
		feed.PATCH("/:page_post_id", handler.FeedUpdate)
	}
}
