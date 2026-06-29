package public

import (
	"api-for-shops-on-facebook-page/configs"
	"api-for-shops-on-facebook-page/routers"

	"github.com/gin-gonic/gin"
)

func Index() {
	r := gin.Default()

	// r.Static("/assets", "./assets")
	// r.GET("/test", func(ctx *gin.Context) {
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// 	ctx.Writer.Header().Set("Content-Type", "image/jpeg")
	// 	ctx.StaticFile("/a.jpg")
	// })
	routers.FacebookAPI(r)

	r.Run(configs.LoadEnv().AppPort)
}
