package auth

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/olegrok/GoHeartRate/protocol"
	"log"
	"net/http"
	"strings"
)

var login, password string

// StartLogin provides minimal console user interface for authorization or registration
func StartLogin(client *http.Client) *http.Response {
	for {
		fmt.Printf("Are you already registred? [y/n]\n")
		var s string
		var errorMsg protocol.ErrorData
		if n, err := fmt.Scanln(&s); err != nil && n > 0 {
			log.Fatalf("input error: %s", err)
		}
		switch strings.ToLower(s) {
		case "no":
			fallthrough
		case "n":
			if !registrationInterface(client, &errorMsg) {
				break
			}
			fallthrough
		case "yes":
			fallthrough
		case "y":
			if res := loginInterface(client); res.StatusCode != http.StatusOK {
				errorMsg = protocol.BytesToErrorData(res.Body)
				break
			} else {
				return res
			}
		default:
			fmt.Println("Error! Try again!")
			continue
		}
		errorHandler(errorMsg)
	}
	return nil
}

func loginInterface(client *http.Client) *http.Response {
	fmt.Print("Enter login: ")
	if _, err := fmt.Scanln(&login); err != nil {
		log.Printf("login error: %s", err)
	}

	fmt.Print("Enter password: ")
	if passByte, err := gopass.GetPasswd(); err != nil {
		log.Fatalf("login error: %s", err)
	} else {
		password = string(passByte)
	}
	res, err := authorization(client, login, password)
	if err != nil {
		log.Fatalf("authorization error: %s", err)
	}
	return res
}

func registrationInterface(client *http.Client, errorMsg *protocol.ErrorData) bool {
	fmt.Print("Enter new login: ")
	if _, err := fmt.Scanln(&login); err != nil {
		log.Fatalf("login error: %s", err)
	}

	fmt.Print("Enter password: ")
	if passByte, err := gopass.GetPasswd(); err != nil || len(passByte) == 0 {
		log.Fatalf("login error: %s", err)
	} else {
		password = string(passByte)
	}
	if len(password) <= 1 {
		log.Fatalf("login error: password is too short\n")
	}
	for rePassword := ""; rePassword != password; {
		fmt.Print("Repeat password: ")
		passByte, err := gopass.GetPasswd()
		if err != nil {
			log.Fatalf("login error: %s", err)
		}
		rePassword = string(passByte)
	}

	res, err := registration(client, login, password)
	if err != nil {
		log.Fatalf("authorization error: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		*errorMsg = protocol.BytesToErrorData(res.Body)
		fmt.Printf("Response ststus code = %d\n", res.StatusCode)
		return false
	}
	fmt.Println("*************************\n*\tSUCCESS!\t*\n*************************")
	return true
}

func errorHandler(err protocol.ErrorData) {
	switch err.ErrorCode {
	case protocol.AlreadyRegistered:
		fallthrough
	case protocol.WrongPassword:
		fallthrough
	case protocol.CalculationError:
		fallthrough
	case protocol.JobTimedOut:
		fallthrough
	case protocol.DatabaseError:
		fallthrough
	default:
		log.Printf("Error code %d: %s\n", err.ErrorCode, err.Error)
	}
}
