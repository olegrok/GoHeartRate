package requests

import (
	"github.com/olegrok/GoHeartRate/protocol"
	"log"
	"net/http"
)

// GetLastResults sends last 10 measurements to user
func GetLastResults(client *http.Client) (*http.Response, error) {
	req, err := CreateRequest(protocol.TransmittedMessage{
		MessageType: protocol.Results,
	})

	if err != nil {
		log.Fatalf("data send error: %s", err)
	}

	return client.Do(req)
}
