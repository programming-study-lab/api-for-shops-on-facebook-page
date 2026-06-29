package models

import "time"

type ContentPost struct {
	Message              string    `json:"message"`
	Link                 string    `json:"link"`
	Published            string    `json:"published"`
	ScheduledPublishTime time.Time `json:"sheduledPublishTime"`
}

type ConversationsModel struct {
}
