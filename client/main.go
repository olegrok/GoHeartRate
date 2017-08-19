package main

import (
	"net/http"
	"log"
	"time"
	"fmt"
	"github.com/olegrok/GoHeartRate/client/auth"
)

const URL = "http://localhost:8080/"

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	login := "login"
	password := "password"

	response, err := auth.Authorization(client, login, password)
	if err != nil {
		log.Printf("authorization error: %s", err)
		return
	}
	fmt.Println(response)
}
