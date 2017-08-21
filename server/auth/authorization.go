package auth

import (
	"fmt"
	"net/http"

	"crypto/md5"
	"github.com/olegrok/GoHeartRate/protocol"
	"github.com/olegrok/GoHeartRate/server/database"
	"github.com/satori/go.uuid"
	"io"
)

func Authorization(w http.ResponseWriter, data protocol.AuthData) {
	fmt.Printf("Authorization:\n login: %s, password: %s\n", data.Login, data.Password)

	ok, session := isRightPassword(data.Login, data.Password)
	fmt.Println(ok, session)
	if ok {
		//go func() {
		database.DB.Create(&session)
		//}()
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
		w.Write(protocol.ErrorDataToBytes(protocol.ErrWrongPassword, protocol.WrongPassword))
	}

}

func isRightPassword(login string, password string) (bool, *database.UserSession) {
	var usr database.User

	if database.DB.Where("username = ?", login).First(&usr).RecordNotFound() {
		return false, nil
	}
	fmt.Println()
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

func SaltPassword(password string) (string, string) {
	h := md5.New()
	salt := fmt.Sprint(uuid.NewV4())
	io.WriteString(h, fmt.Sprintf("%s::%s", password, salt))
	fmt.Printf("%s\n%x\n", salt, h.Sum(nil))
	return fmt.Sprintf("%x", h.Sum(nil)), salt
}

func IsValidPassword(inputPassword string, salt string, realPassword string) bool {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%s::%s", inputPassword, salt))
	fmt.Printf("%s\n%x\n%s\n", salt, h.Sum(nil), realPassword)
	return realPassword == fmt.Sprintf("%x", h.Sum(nil))
}
