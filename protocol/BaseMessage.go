package protocol

import "encoding/json"

type TransmittedMessage struct {
	MessageType string		`json:"type"`
	Data interface{}		`json:"data"`
}

type ReceivedMessage struct {
	MessageType string		`json:"type"`
	Data json.RawMessage	`json:"data"`
}
