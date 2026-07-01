package facebookmodel

type MessageModel struct {
	RecipientId   string `json:"recipientId"`
	MessagingType string `json:"messagingType"`
	MessageText   string `json:"messageText"`
	MediaType     string `json:"mediaType"`
}

type FacebookId struct {
	Id string `json:"id"`
}

type FacebookMessageText struct {
	Text string `json:"text"`
}

type FacebookMessageModel struct {
	Recipient     FacebookId          `json:"recipient"`
	MessagingType string              `json:"RESPONSE"`
	Message       FacebookMessageText `json:"message"`
	AccessToken   string              `json:"accessToken"`
}

type FacebookMessageAttachment struct {
	AttachmentType string            `json:"type"`
	Payload        AttachmentPayload `json:"payload"`
}

type AttachmentPayload struct {
	IsReusable AttachmentIsReusable `json:"isReusable"`
}

type AttachmentIsReusable struct {
	IsReusable bool `json:"isReusable"`
}
