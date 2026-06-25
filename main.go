package main

import (
	"api-for-shops-on-facebook-page/configs"
	"api-for-shops-on-facebook-page/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.FacebookAPI(r)

	r.Run(configs.LoadEnv().Port_Host)
}
