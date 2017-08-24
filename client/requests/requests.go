package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/olegrok/GoHeartRate/protocol"
)

const url = "http://localhost:8080/"

// CreateRequest creates request object to sending to server
func CreateRequest(msg protocol.TransmittedMessage) (*http.Request, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
		return nil, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		log.Fatalf("new request error: %s", err)
		return nil, err
	}
	return request, nil
}
