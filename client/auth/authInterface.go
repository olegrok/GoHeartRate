package auth

import (
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/olegrok/GoHeartRate/protocol"
	"log"
	"net/http"
	"strings"
)

func StartLogin(client *http.Client) (*http.Response, error) {
	var login, password string
	for {
		fmt.Printf("Are you already registred? [y/n]\n")
		var s string
		var errorMsg protocol.ErrorData
		fmt.Scanln(&s)
		switch strings.ToLower(s) {
		case "no":
			fallthrough
		case "n":
			fmt.Print("Enter new login: ")
			if _, err := fmt.Scanln(&login); err != nil {
				log.Printf("login error: %s", err)
			}

			fmt.Print("Enter password: ")
			passByte, err := gopass.GetPasswd()
			if err != nil {
				log.Fatalf("login error: %s", err)
			}
			password = string(passByte)

			for rePassword := ""; rePassword != password; {
				fmt.Print("Repeat password: ")
				passByte, err := gopass.GetPasswd()
				if err != nil {
					log.Fatalf("login error: %s", err)
				}
				rePassword = string(passByte)
			}

			res, err := Registration(client, login, password)
			if err != nil {
				log.Fatalf("authorization error: %s", err)
			}
			if res.StatusCode != http.StatusOK {
				errorMsg = protocol.BytesToErrorData(res.Body)
				fmt.Printf("Response ststus code = %d\n", res.StatusCode)
				break
			}
			fmt.Println("*************************\n*\tSUCCESS!\t*\n*************************")
			fallthrough
		case "yes":
			fallthrough
		case "y":
			fmt.Print("Enter login: ")
			if _, err := fmt.Scanln(&login); err != nil {
				log.Printf("login error: %s", err)
			}

			fmt.Print("Enter password: ")
			passByte, err := gopass.GetPasswd()
			if err != nil {
				log.Fatalf("login error: %s", err)
			}
			password = string(passByte)

			res, err := Authorization(client, login, password)
			if err != nil {
				log.Fatalf("authorization error: %s", err)
			}
			if res.StatusCode != http.StatusOK {
				errorMsg = protocol.BytesToErrorData(res.Body)
				fmt.Printf("Response ststus code = %d\n", res.StatusCode)
				break
			}
			return res, err

		default:
			fmt.Println("Error! Try again!")
			continue
		}

		switch errorMsg.ErrorCode {
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
			fmt.Printf("Error code %d: %s\n", errorMsg.ErrorCode, errorMsg.Error)
		}
	}
	return nil, errors.New("Unknown login error")
}
