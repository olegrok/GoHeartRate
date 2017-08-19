package auth

import (
	"github.com/olegrok/GoHeartRate/protocol"
	"net/http"
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
