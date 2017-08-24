package auth

import (
	"fmt"
	"net/http"

	"crypto/md5"
	"github.com/olegrok/GoHeartRate/protocol"
	"github.com/olegrok/GoHeartRate/server/database"
	"github.com/satori/go.uuid"
	"io"
	"log"
)

// Authorization send request to auth
func Authorization(w http.ResponseWriter, data protocol.AuthData) {
	fmt.Printf("Authorization:\n login: %s, password: %s\n", data.Login, data.Password)

	ok, session := isRightPassword(data.Login, data.Password)
	fmt.Println(ok, session)
	if ok {
		database.DB.Create(&session)
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: session.Token,
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "uid",
			Value: fmt.Sprint(session.UserID),
		})
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		if _, err := w.Write(protocol.ErrorDataToBytes(protocol.ErrWrongPassword, protocol.WrongPassword)); err != nil {
			log.Printf("response writer error: %s", err)
		}
	}

}

//isRightPassword check login and password and returns boolean value and user session struct if user is valid
func isRightPassword(login string, password string) (bool, *database.UserSession) {
	var usr database.User

	if database.DB.Where("username = ?", login).First(&usr).RecordNotFound() {
		return false, nil
	}
	fmt.Printf("usr pass: %s\n", usr.Password)
	if IsValidPassword(password, usr.Salt, usr.Password) {
		session := database.UserSession{
			UserID: usr.ID,
			Token:  uuid.NewV4().String(),
		}
		return true, &session
	}
	return false, nil
}

// SaltPassword returns password hash and salt
func SaltPassword(password string) (string, string) {
	h := md5.New()
	salt := fmt.Sprint(uuid.NewV4())
	if _, err := io.WriteString(h, fmt.Sprintf("%s::%s", password, salt)); err != nil {
		log.Printf("salt password error: %s", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), salt
}

// IsValidPassword checks input hash with database value
func IsValidPassword(inputPassword string, salt string, realPassword string) bool {
	h := md5.New()
	if _, err := io.WriteString(h, fmt.Sprintf("%s::%s", inputPassword, salt)); err != nil {
		log.Printf("check password validity error: %s", err)
	}
	return realPassword == fmt.Sprintf("%x", h.Sum(nil))
}
