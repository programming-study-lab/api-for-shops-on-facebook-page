package facebookcontroller

import (
	"api-for-shops-on-facebook-page/configs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huandu/facebook/v2"
)

func FacebookPageGetInfo(ctx *gin.Context) {
	fbConfig := configs.ConnectFacebookGraphApi()

	res, err := facebook.Get("/me", facebook.Params{
		"fields":       "id,name",
		"access_token": fbConfig.AccessToken,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadGateway,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusOK,
			gin.H{
				"status":  true,
				"message": "success",
				"data": gin.H{
					"id":   res.GetField("id"),
					"name": res.GetField("name"),
				},
			},
		)
	}
}

func FacebookPageGetConversations(ctx *gin.Context) {
	fbConfig := configs.ConnectFacebookGraphApi()

	path := "/" + fbConfig.FacebookPageId
	path += "/conversations"
	res, err := facebook.Get(path, facebook.Params{
		// "fields": "messages,from,created_time",
		// "fields":       "participants,name,id,created_time,from,to,reply_to",
		"fields":       "id,is_owner,participants,name,username,updated_time",
		"access_token": fbConfig.AccessToken,
	})

	if err == nil {
		ctx.AbortWithStatusJSON(
			http.StatusOK,
			gin.H{
				"status":  true,
				"message": "success",
				"data":    res.GetField("data"),
			},
		)
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
	}
}

func FacebookPageGetMessages(ctx *gin.Context) {
	fbConfig := configs.ConnectFacebookGraphApi()

	path := "/" + fbConfig.FacebookPageId
	path += "/conversations"

	conversation_id := ctx.Param("conversation_id")

	message_path := fmt.Sprintf("/%s", conversation_id)
	message_path += "/messages"

	res, err := facebook.Get(message_path, facebook.Params{
		// "fields":       "message,from,created_time",
		"fields":       "id,created_time,from,to,message,reply_to",
		"access_token": fbConfig.AccessToken,
	})

	if err == nil {
		ctx.AbortWithStatusJSON(
			http.StatusOK,
			gin.H{
				"status":  true,
				"message": "success",
				"data":    res.GetField("data"),
			},
		)
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
	}

}

func FacebookPageSendMessage(ctx *gin.Context) {

}

func GetWebhookController(ctx *gin.Context) {
	mode := ctx.Query("hub.mode")
	webhook_token := ctx.Query("hub.verify_token")
	challenge := ctx.Query("hub.challenge")

	fbConfig := configs.ConnectFacebookGraphApi()

	if mode == "subscribe" && webhook_token == fbConfig.FacebookWebhookToken {
		ctx.String(
			http.StatusOK,
			challenge,
		)
		return
	}

	ctx.Status(
		http.StatusForbidden,
	)
}

func PostWebhookController(ctx *gin.Context) {
	var body map[string]interface{}

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.Status(
			http.StatusBadGateway,
		)

		// fmt.Printf("\n[facebook] error = %s\n", err.Error())
		return
	}

	// ctx.AbortWithStatusJSON(
	// 	http.StatusOK,
	// 	gin.H{
	// 		"status":  true,
	// 		"message": "EVENT_RECEVIVED",
	// 		"data":    &body,
	// 	},
	// )
	// d := []byte

	// data, _ := json.Marshal(body)

	// fmt.Printf("\n%s\n\n", string(data))

	fmt.Printf("\n[Webhooks] ---\n\n")

	ctx.String(
		http.StatusOK,
		"EVENT_RECEVIVED",
	)

	// fmt.Printf("\n[facebook] EVENT_RECEVIVED = %s\n", body)
}
