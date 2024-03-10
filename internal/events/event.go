package events

import (
	"encoding/json"
	"errors"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

const (
	EventSendMessage = "send_message"
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}
