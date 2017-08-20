package auth

import (
	"net/http"

	"github.com/olegrok/GoHeartRate/protocol"
)

func Registration(w http.ResponseWriter, data protocol.AuthData) {
	if IsLoginNew(data.Login) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(protocol.ErrorDataToBytes(protocol.ErrAlreadyRegistered, protocol.AlreadyRegistered))
	}
}

func IsLoginNew(login string) bool {
	// todo Check in DataBase
	return true
}
