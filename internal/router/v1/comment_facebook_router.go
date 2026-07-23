package v1

import (
	"api-for-shops-on-facebook-page/internal/module/comment_facebook/delivery/http"

	"github.com/gin-gonic/gin"
)

func CommentFacebookRouter(groupV1 *gin.RouterGroup, handler *http.CommentFacebookHttp) {
	comment := groupV1.Group("/comments")
	{
		// comment.GET("/:feedId", handler.CommentList)
		comment.GET("/", handler.CommentList)
		comment.POST("/")
		comment.DELETE("/:comment_id")
	}
}
