package test

import (
	"api-for-shops-on-facebook-page/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huandu/facebook/v2"
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

func OnTest(ctx *gin.Context) {
	fbConfig := configs.ConnectFacebookGraphApi()

	res, err := facebook.Get("/me", facebook.Params{
		"fields":       "id,name",
		"access_token": fbConfig.AccessToken,
	})

	// res, err := fb.Get("/"+fbConfig.FacebookPageId, fb.Params{
	// 	"fields":       "id,name",
	// 	"access_token": fbConfig.AccessToken,
	// })

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err,
			},
		)
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusOK,
			gin.H{
				"status":  true,
				"message": "success",
				// "data":    fbConfig.VersionGraph,
				"res": res,
			},
		)
	}

}
