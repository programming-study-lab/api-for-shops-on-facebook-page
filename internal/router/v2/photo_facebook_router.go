package v1

import (
	"api-for-shops-on-facebook-page/internal/module/photo_facebook/delivery/http"

	"github.com/gin-gonic/gin"
)

func PhotoFacebookRouter(group *gin.RouterGroup, handler *http.PhotoFacebookHttp) {
	photo := group.Group("/photos")
	{
		photo.POST("/", handler.PhotoCreate)
	}
}
