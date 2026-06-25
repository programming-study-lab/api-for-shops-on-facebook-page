package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FacebookPage(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "success",
			"data":    "",
		},
	)
}
