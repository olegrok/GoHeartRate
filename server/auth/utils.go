package auth

import (
	"encoding/json"
	"github.com/olegrok/GoHeartRate/protocol"
	"log"
)

// GetAuthMessage unmarshals raw json with authentication information to special structure
func GetAuthMessage(data json.RawMessage) (*protocol.AuthData, error) {
	var msg protocol.AuthData
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Printf("marshal message error: %s", err)
		return nil, err
	}
	return &msg, nil
}
