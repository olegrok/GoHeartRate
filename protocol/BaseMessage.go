package protocol

import (
	"encoding/json"
	"strings"
)

type messageType int

const (
	// Auth is a code that uses to sending with authorization messages
	Auth messageType = iota + 1
	// Register is a code that uses to sending with registration messages
	Register
	// Data is a code that uses to sending messages with results of measurements
	Data
	// Results is a code to request last 10 results of measurements
	Results
	// Unknown is a special code for unknown messages
	Unknown
)

// TransmittedMessage is a structure to send messages of different types and their code
type TransmittedMessage struct {
	MessageType messageType `json:"type"`
	Data        interface{} `json:"data"`
}

// ReceivedMessage structure that contains data in raw to future marshaling
type ReceivedMessage struct {
	MessageType messageType     `json:"type"`
	Data        json.RawMessage `json:"data"`
}

func (m *messageType) UnmarshalJSON(b []byte) error {
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
	case "results":
		*m = Results
	}

	return nil
}

func (m messageType) MarshalJSON() ([]byte, error) {
	var s string
	switch m {
	default:
		s = "unknown"
	case Auth:
		s = "auth"
	case Register:
		s = "register"
	case Data:
		s = "math"
	case Results:
		s = "results"
	}

	return json.Marshal(s)
}
