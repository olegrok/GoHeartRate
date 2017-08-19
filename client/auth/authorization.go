package auth

import (
	"github.com/olegrok/GoHeartRate/client/requests"
	"github.com/olegrok/GoHeartRate/protocol"
	"log"
	"net/http"
)

func Authorization(client *http.Client, login string, password string) (*http.Response, error) {
	msg := protocol.TransmittedMessage{
		MessageType:     "auth",
		MessageTypeCode: protocol.Auth,
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
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("response error: %s", err)
		return nil, err
	}

	return response, nil
}
