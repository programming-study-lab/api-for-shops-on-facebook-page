package routers

import (
	controller "api-for-shops-on-facebook-page/controllers/facebook-controller"
	"api-for-shops-on-facebook-page/test"

	"github.com/gin-gonic/gin"
)

func FacebookAPI(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/test", test.CheckApi)
		api.GET("content", controller.FacebookPage)
	}

}
