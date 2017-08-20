package auth

import (
	"fmt"
	"net/http"

	"github.com/olegrok/GoHeartRate/protocol"
	"github.com/olegrok/GoHeartRate/server/database"
	"github.com/satori/go.uuid"
)

func Authorization(w http.ResponseWriter, data protocol.AuthData) {
	fmt.Printf("Authorization:\n login: %s, password: %s\n", data.Login, data.Password)

	ok, session := isRightPassword(data.Login, data.Password)
	fmt.Println(ok, session)
	if ok {
		go func() {
			database.DB.Create(&session)
		}()
		cookie := http.Cookie{
			Name:     "token",
			Value:    session.Token,
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

func isRightPassword(login string, password string) (bool, *database.UserSession) {
	var usr database.User
	if database.DB.Where("username = ? AND password = ?", login, password).First(&usr).RecordNotFound() {
		return false, nil
	}
	session := database.UserSession{
		UserID: usr.ID,
		Token:  uuid.NewV4().String(),
	}
	return true, &session
}
