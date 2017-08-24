package auth

import (
	"github.com/olegrok/GoHeartRate/protocol"
	"github.com/olegrok/GoHeartRate/server/database"
	"log"
	"net/http"
	"time"
)

// Registration create new user if it does not exists or returns error if user already registered
func Registration(w http.ResponseWriter, data protocol.AuthData) {
	if IsLoginNew(data.Login) {
		pass, salt := SaltPassword(data.Password)
		database.DB.Create(&database.User{
			Username:  data.Login,
			Password:  pass,
			Salt:      salt,
			CreatedAt: time.Now(),
		})
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		if _, err := w.Write(protocol.ErrorDataToBytes(protocol.ErrAlreadyRegistered, protocol.AlreadyRegistered)); err != nil {
			log.Printf("registration error: %s", err)
		}
	}
}

// IsLoginNew checks the uniqueness of login. Returns "true" if login is new
func IsLoginNew(login string) bool {
	return database.DB.Where("username = ?", login).First(&database.User{}).RecordNotFound()
}
