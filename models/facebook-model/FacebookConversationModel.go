package facebookmodel

import (
	"mime/multipart"

	"github.com/huandu/facebook/v2"
)

type FacebookMessageModelBase struct {
	MessageRequestModel                MessageRequestModel                `json:"messageRequestModel"`
	FacebookMessageModel               FacebookMessageModel               `json:"FacebookMessageModel"`
	FacebookSendMessageAttachmentModel FacebookSendMessageAttachmentModel `json:"FacebookMessageAttachmen"`
}

type MessageRequestModel struct {
	RecipientId   string `json:"recipientId" form:"recipientId"`
	MessagingType string `json:"messagingType"`
	MessageText   string `json:"messageText"`
	MediaType     string `json:"mediaType" form:"mediaType"`
	Message       string `form:"message"`
	Recipient     string `form:"recipient"`
	// Filedata  *multipart.FileHeader `json:"filedata"`
	Filedata multipart.FileHeader `form:"filedata"`
}

type FacebookId struct {
	Id string `json:"id"`
}

type FacebookMessageText struct {
	Text string `json:"text"`
}

// ส่งข้อมูล File รูปภาพ จาก Header: form-data
type FacebookMessageModel struct {
	Recipient     FacebookId          `json:"recipient"`
	MessagingType string              `json:"messageingType"`
	Message       FacebookMessageText `json:"message"`
	AccessToken   string              `json:"accessToken"`
}

type FacebookSendMessageAttachmentModel struct {
	Recipient string               `json:"recipient"`
	Message   string               `json:"message"`
	Filedata  *facebook.BinaryData `json:"filedata"`
}

type Recipient struct {
	Id string `json:"id"`
}

type Message struct {
	Attachment Attachment `json:"attachment"`
}

type Filedata struct {
}

type Attachment struct {
	AttachmentType string            `json:"type"`
	Payload        AttachmentPayload `json:"payload"`
}

type AttachmentPayload struct {
	IsReusable bool `json:"is_reusable"`
}

type AttachmentIsReusable struct {
	IsReusable bool `json:"isReusable"`
}
