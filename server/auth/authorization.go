package auth

import (
	"net/http"
	"github.com/olegrok/GoHeartRate/protocol"
	//"time"
	"fmt"
	"encoding/json"
	"log"
)

func Authorization(w http.ResponseWriter, r *http.Request, data protocol.AuthData){
	fmt.Printf("Authorization:\n login: %s, password: %s\n", data.Login, data.Password)

	ok, cookieValue := isRightPassword(data.Login, data.Password)
	if ok {
		cookie := http.Cookie {
			Name: "access_key",
			Value: cookieValue,
			MaxAge: 86400,
			Secure: true,
			HttpOnly: true,
			//Raw: "access_key="+cookieValue,
			//[]string{"test=tcookie"},
		}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
	} else {
		errorMsg := protocol.ErrorData{
			Message: "Login is not found or password is wrong",
			MessageCode:protocol.WrongPassword,
		}
		data, err := json.Marshal(errorMsg)
		if err != nil {
			log.Fatalf("marshaling error: %s", err)
		}
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(data))
	}

}

func IsAuthorized(w http.ResponseWriter, r *http.Request, data protocol.AuthData) (bool) {
	//todo Check in Database
	return true
}

func isRightPassword(login string, password string) (bool, string) {
	//todo Check in DataBase
	return true, "COOKIE"
}