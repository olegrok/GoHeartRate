package auth

import (
	"github.com/olegrok/GoHeartRate/protocol"
	"encoding/json"
	"log"
	"bytes"
	"net/http"
)

const URL = "http://localhost:8080/"

func Authorization(client *http.Client, login string, password string) (*http.Response, error){
	msg := protocol.TransmittedMessage{
		"auth",
		protocol.AuthData {
			login,
			password,
		},
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
		return nil, err
	}
	jsonStr := []byte(data)
	request, err := http.NewRequest("POST", URL, bytes.NewReader(jsonStr))
	if err != nil {
		log.Fatalf("new request error: %s", err)
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("response error: %s", err)
		return nil, err
	}

	return response, nil
}
