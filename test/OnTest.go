package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckApi(ctx *gin.Context) {

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "success",
			"data":    "[{}]",
		},
	)

}
