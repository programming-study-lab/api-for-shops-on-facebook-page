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
		api.GET("/test", test.CheckApi)                      // ทดสอบว่าระบบ API สามารถใช้งานได้หรือไม่
		api.GET("/get-info", controller.FacebookPageGetInfo) // ดึงข้อมูลของเพจ

		api.GET("/get-conversations", controller.FacebookPageGetConversations)        // ดึงรายการสนทนาของเพจ
		api.GET("/get-messages/:conversation_id", controller.FacebookPageGetMessages) // ดึงข้อความแชทจากเพจ

		// api.Static("/images", "./uploads")

		api.GET("/get-image/:image_name", controllers.GetImage) // แสดงรูปภาพ
		api.POST("/upload-image", controllers.UploadImage)      // อัพโหลดรูปภาพ

		api.GET("/webhook", controller.GetWebhookController)   // ตรวจสอบยืนยัน web hooks
		api.POST("/webhook", controller.PostWebhookController) // รับข้อความจาก web hooks เช่น การแจ้งเตือน

		api.POST("/send-message", controller.FacebookPageSendMessage) // ส่งข้อความ และ ส่งรูปภาพ ไปยัง ลูกค้าของเพจ
	}

}
