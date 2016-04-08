package minebot

import "sync/atomic"

var msgID int32 = 0

type slackMessage struct {
	ID      int32  `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func NewSlackMessage(channel string, text string) slackMessage {
	nextID := atomic.AddInt32(&msgID, 1)
	return slackMessage{nextID, "message", channel, text}
}
