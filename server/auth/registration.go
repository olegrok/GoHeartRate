package auth

import (
	"github.com/olegrok/GoHeartRate/protocol"
	"net/http"
	"encoding/json"
	"log"
)

func Registration(w http.ResponseWriter, r *http.Request, data protocol.AuthData) {
	if IsLoginNew(data.Login) {
		w.WriteHeader(http.StatusOK)
	} else {
		errorMsg := protocol.ErrorData{
			Message: "Username is already registered",
			MessageCode: protocol.AlreadyRegistered,
		}
		data, err := json.Marshal(errorMsg)
		if err != nil {
			log.Fatalf("marshaling error: %s", err)
		}
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(data))
	}
}

func IsLoginNew(login string) (bool) {
	// todo Check in DataBase
	return true
}
