package requests

import (
	"bytes"
	"encoding/json"
	"github.com/olegrok/GoHeartRate/protocol"
	"log"
	"net/http"
)

func CreateRequest(msg protocol.TransmittedMessage) (*http.Request, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
		return nil, err
	}
	request, err := http.NewRequest("POST", protocol.URL, bytes.NewReader([]byte(data)))
	if err != nil {
		log.Fatalf("new request error: %s", err)
		return nil, err
	}
	return request, nil
}