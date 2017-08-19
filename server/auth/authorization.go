package auth

import (
	"fmt"
	"github.com/olegrok/GoHeartRate/protocol"
	"net/http"
)

func Authorization(w http.ResponseWriter, data protocol.AuthData) {
	fmt.Printf("Authorization:\n login: %s, password: %s\n", data.Login, data.Password)

	ok, cookieValue := isRightPassword(data.Login, data.Password)
	if ok {
		cookie := http.Cookie{
			Name:     "access_key",
			Value:    cookieValue,
			MaxAge:   86400,
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(protocol.ErrorDataToBytes(protocol.ErrWrongPassword, protocol.WrongPassword))
	}

}

func IsAuthorized(w http.ResponseWriter, r *http.Request, data protocol.AuthData) bool {
	//todo Check in Database
	return true
}

func isRightPassword(login string, password string) (bool, string) {
	//todo Check in DataBase
	return true, "COOKIE"
}
