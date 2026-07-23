package http

import (
	"api-for-shops-on-facebook-page/internal/module/conversation_facebook/domain"
	facebookmodel "api-for-shops-on-facebook-page/models/facebook-model"
	facebookservices "api-for-shops-on-facebook-page/services/facebook-service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huandu/facebook/v2"
)

type ConversationFacebookHttp struct {
	usecase domain.ConversationFacebookUsecase
}

func NewConversationFacebookHttp(usecase domain.ConversationFacebookUsecase) *ConversationFacebookHttp {
	return &ConversationFacebookHttp{
		usecase: usecase,
	}
}

// ดึงข้อมูลการสนทนาของเพจ Facebook
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

// ดึงข้อความแชทของเพจแชทกับแชทกับลูกค้า
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
		"fields": "id,created_time,from,to,message,reply_to,attachments",
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

// ส่งข้อความแชทไปยังลูกค้า
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

	messageModel := facebookmodel.MessageRequestModel{}
	if err := ctx.ShouldBind(&messageModel); err != nil {

		log.Printf("\n[json] %v,\n", messageModel)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		fmt.Printf("\ntest\n")
		return
	}

	facebookMessageModel := facebookmodel.FacebookMessageModelBase{}

	if messageModel.MediaType == "image" || messageModel.MediaType == "video" {

		// ทำให้เป็น string ก่อนส่งไปยัง facebook
		recipientId, err := json.Marshal(facebookmodel.Recipient{
			Id: messageModel.RecipientId,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"status":  false,
					"message": "เกิดความผิดปกติเกี่ยวกับข้อมูลใน Recipient",
					"data":    err.Error(),
				},
			)
			return
		}

		// ทำให้เป็น string ก่อนส่งไปยัง facebook
		messageAttachment, err := json.Marshal(facebookmodel.Message{
			Attachment: facebookmodel.Attachment{
				AttachmentType: messageModel.MediaType,
				Payload: facebookmodel.AttachmentPayload{
					IsReusable: true,
				},
			},
		})

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"status":  false,
					"message": "เกิดความผิดปกติเกี่ยวกับข้อมูลใน Message",
					"data":    err.Error(),
				},
			)
			return
		}

		file, err := messageModel.Filedata.Open()
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"status":  false,
					"message": "เกิดความผิดปกติเกี่ยวกับไฟล์ที่ส่งมายัง Server กลาง",
					"data":    err.Error(),
				},
			)
			return
		}

		facebookMessageModel = facebookmodel.FacebookMessageModelBase{
			FacebookSendMessageAttachmentModel: facebookmodel.FacebookSendMessageAttachmentModel{
				Recipient: string(recipientId),
				Message:   string(messageAttachment),
				Filedata:  facebook.Data(messageModel.Filedata.Filename, file),
			},
		}
	} else {
		facebookMessageModel = facebookmodel.FacebookMessageModelBase{
			FacebookMessageModel: facebookmodel.FacebookMessageModel{
				Recipient: facebookmodel.FacebookId{
					Id: messageModel.RecipientId,
				},
				MessagingType: "RESPONSE",
				Message: facebookmodel.FacebookMessageText{
					Text: messageModel.MessageText,
				},
			},
		}
	}

	if messageModel.MediaType == "image" || messageModel.MediaType == "video" {

		res, err := session.Post("/me/messages", facebook.Params{
			"recipient": facebookMessageModel.FacebookSendMessageAttachmentModel.Recipient,
			"message":   facebookMessageModel.FacebookSendMessageAttachmentModel.Message,
			"filedata":  facebookMessageModel.FacebookSendMessageAttachmentModel.Filedata,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"status":  false,
					"message": "เกิดความผิดปกติขณะส่งข้อมูลไปยัง Facebook",
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
					"recipientId":   res.GetField("recipient_id"),
					"messageId":     res.GetField("message_id"),
					"attachment_id": res.GetField("attachment_id"),
				},
			},
		)
	} else {

		res, err := session.Post("/me/messages", facebook.MakeParams(facebookMessageModel.FacebookMessageModel))

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"status":  false,
					"message": "เกิดความผิดปกติในการส่งข้อความไปยัง Facebook",
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

}
