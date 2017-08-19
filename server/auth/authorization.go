package auth

import (
	"net/http"
	"github.com/olegrok/GoHeartRate/protocol"
	"time"
	"fmt"
)

func Authorization(w http.ResponseWriter, r *http.Request, data protocol.AuthData){
	fmt.Printf("Authorization:\n login: %s, password: %s\n", data.Login, data.Password)
	cookie := http.Cookie {
		"Name",
		"Value",
		"/",
		"",
		time.Now().AddDate(0, 0, 1),
		time.Now().AddDate(0, 0, 1).Format(time.UnixDate),
		86400,
		true,
		true,
		"test=tcookie",
		[]string{"test=tcookie"},
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusAccepted)
}
