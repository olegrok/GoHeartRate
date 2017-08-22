package math

import (
	"log"
	"net/http"

	"github.com/olegrok/GoHeartRate/client/requests"
	"github.com/olegrok/GoHeartRate/protocol"
)

func Transmit(client *http.Client, array []float64) (*http.Response, error) {
	dataMsg := protocol.MathData{
		DataArray: array,
		TimeArray: []float64{1, 2, 3},
	}

	req, err := requests.CreateRequest(protocol.TransmittedMessage{
		MessageType: protocol.Data,
		Data:        dataMsg,
	})

	if err != nil {
		log.Fatalf("data send error: %s", err)
	}

	return client.Do(req)
}
