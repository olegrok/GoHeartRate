package protocol

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MessageType int

const (
	Auth MessageType = iota + 1
	Register
	Data
	Unknown
)

type TransmittedMessage struct {
	MessageType MessageType `json:"type"`
	Data        interface{} `json:"data"`
}

type ReceivedMessage struct {
	MessageType     MessageType     `json:"type"`
	MessageTypeCode int             `json:"type_code"`
	Data            json.RawMessage `json:"data"`
}

func (m *MessageType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*m = Unknown
	case "auth":
		*m = Auth
	case "register":
		*m = Register
	case "math":
		*m = Data
	}

	return nil
}

func (a MessageType) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Auth:
		s = "auth"
	case Register:
		s = "register"
	case Data:
		s = "math"
	}

	byte, _ := json.Marshal(s)
	fmt.Printf("%s\n", byte)
	return json.Marshal(s)
}
