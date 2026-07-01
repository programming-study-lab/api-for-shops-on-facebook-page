package routers

import (
	"api-for-shops-on-facebook-page/controllers"
	controller "api-for-shops-on-facebook-page/controllers/facebook-controller"
	"api-for-shops-on-facebook-page/test"

	"github.com/gin-gonic/gin"
)

func FacebookAPI(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/test", test.CheckApi)
		api.GET("/get-info", controller.FacebookPageGetInfo)

		api.GET("/get-conversations", controller.FacebookPageGetConversations)
		api.GET("/get-messages/:conversation_id", controller.FacebookPageGetMessages)

		api.Static("/images", "./uploads")

		api.GET("/get-image/:image_name", controllers.GetImage)
		api.POST("/upload-image", controllers.UploadImage)

		api.GET("/webhook", controller.GetWebhookController)
		api.POST("/webhook", controller.PostWebhookController)

		api.POST("/send-message", controller.FacebookPageSendMessage)
	}

}
