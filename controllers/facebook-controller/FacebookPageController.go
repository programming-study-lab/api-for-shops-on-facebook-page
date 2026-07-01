package facebookcontroller

import (
	"api-for-shops-on-facebook-page/configs"
	facebookmodel "api-for-shops-on-facebook-page/models/facebook-model"
	facebookservices "api-for-shops-on-facebook-page/services/facebook-service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huandu/facebook/v2"
)

func FacebookPageGetInfo(ctx *gin.Context) {
	session, err := facebookservices.FacebookInit()

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err,
			},
		)
		return
	}

	res, err := session.Get("/me", facebook.Params{
		"fields": "id,name",
	})

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err,
			},
		)
		return
	}

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

func FacebookPageGetConversations(ctx *gin.Context) {
	session, err := facebookservices.FacebookInit()

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		return
	}

	res, err := session.Get("/me/conversations", facebook.Params{
		"fields": "id,is_owner,participants,name,username,updated_time",
	})

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "success",
			"data":    res.GetField("data"),
		},
	)
}

func FacebookPageGetMessages(ctx *gin.Context) {
	session, err := facebookservices.FacebookInit()

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		return
	}

	conversation_id := ctx.Param("conversation_id")

	path := "/" + conversation_id + "/messages"

	res, err := session.Get(path, facebook.Params{
		"fields": "id,created_time,from,to,message,reply_to",
		// "fields": "id,created_time,from,to,message,reply_to",
	})

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "success",
			"data":    res.GetField("data"),
		},
	)
}

func FacebookPageSendMessage(ctx *gin.Context) {

	session, err := facebookservices.FacebookInit()

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		return
	}

	messageModel := facebookmodel.MessageModel{}
	if err := ctx.ShouldBindJSON(&messageModel); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		return
	}

	facebookMessageModel := make(map[string]interface{})

	if messageModel.MediaType == "image" {
		facebookMessageModel = facebookmodel.FacebookMessageAttachment{
			AttachmentType: messageModel.MediaType,
			Payload: facebookmodel.AttachmentPayload{
				facebookmodel.AttachmentIsReusable{
					IsReusable: true,
				},
			},
		}
	} else if messageModel.MediaType == "video" {

	} else {
		&facebookMessageModel = facebookmodel.FacebookMessageModel{
			Recipient: facebookmodel.FacebookId{
				Id: messageModel.RecipientId,
			},
			MessagingType: "RESPONSE",
			Message: facebookmodel.FacebookMessageText{
				Text: messageModel.MessageText,
			},
		}

	}

	// facebookMessageModel := facebookmodel.FacebookMessageModel{
	// 	Recipient: facebookmodel.FacebookId{
	// 		Id: messageModel.RecipientId,
	// 	},
	// 	MessagingType: "RESPONSE",
	// 	Message: facebookmodel.FacebookMessageText{
	// 		Text: messageModel.MessageText,
	// 	},
	// }

	res, err := session.Post("/me/messages", facebook.MakeParams(facebookMessageModel))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "success",
			"data": gin.H{
				"recipientId": res.GetField("recipient_id"),
				"messageId":   res.GetField("message_id"),
			},
		},
	)

}

func GetWebhookController(ctx *gin.Context) {
	mode := ctx.Query("hub.mode")
	webhook_token := ctx.Query("hub.verify_token")
	challenge := ctx.Query("hub.challenge")

	fbConfig := configs.FacebookConfig()

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
		return
	}

	log.Printf("\n[Webhooks] ---\n\n")

	ctx.String(
		http.StatusOK,
		"EVENT_RECEVIVED",
	)

}
