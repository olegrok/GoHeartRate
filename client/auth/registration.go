package auth

import (
	"log"
	"net/http"

	"github.com/olegrok/GoHeartRate/client/requests"
	"github.com/olegrok/GoHeartRate/protocol"
)

func registration(client *http.Client, login string, password string) (*http.Response, error) {
	msg := protocol.TransmittedMessage{
		MessageType: protocol.Register,
		Data: protocol.AuthData{
			Login:    login,
			Password: password,
		},
	}

	request, err := requests.CreateRequest(msg)
	if err != nil {
		log.Fatalf("response error: %s", err)
		return nil, err
	}

	return client.Do(request)
}
