package auth

import (
	"github.com/olegrok/GoHeartRate/protocol"
	"github.com/olegrok/GoHeartRate/server/database"
	"net/http"
	"time"
)

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
		w.Write(protocol.ErrorDataToBytes(protocol.ErrAlreadyRegistered, protocol.AlreadyRegistered))
	}
}

func IsLoginNew(login string) bool {
	return database.DB.Where("username = ?", login).First(&database.User{}).RecordNotFound()
}
