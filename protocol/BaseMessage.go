package protocol

import "encoding/json"

const (
	Auth = iota + 1
	Register
	Data
)

type TransmittedMessage struct {
	MessageType string		`json:"type"`
	MessageTypeCode int		`json:"type_code"`
	Data interface{}		`json:"data"`
}

type ReceivedMessage struct {
	MessageType string		`json:"type"`
	MessageTypeCode int		`json:"type_code"`
	Data json.RawMessage	`json:"data"`
}
