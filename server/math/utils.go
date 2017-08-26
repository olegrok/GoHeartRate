package math

import (
	"encoding/json"
	"fmt"
	"github.com/olegrok/GoHeartRate/protocol"
	"github.com/olegrok/GoHeartRate/server/database"
	"io"
	"log"
)

// ResultHandler unmarshals raw data, calculate result and send the response to client. Returns error.
func ResultHandler(RawData json.RawMessage, uid *string, w io.Writer) error {
	var msg protocol.MathData
	if err := json.Unmarshal(RawData, &msg); err != nil {
		log.Printf("marshal message error: %s", err)
		return err
	}
	result := Calculate(msg.DataArray)
	if err := database.SaveResult(*uid, result); err != nil {
		return err
	}

	if _, err := io.WriteString(w, fmt.Sprint(result)); err != nil {
		log.Printf("result response error: %s", err)
		return err
	}
	return nil
}
