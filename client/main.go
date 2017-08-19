package main

import (
	"net/http"
	"log"
	"time"
	"fmt"
	"github.com/olegrok/GoHeartRate/client/auth"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	//response, err := auth.StartLogin(client)
	response, err := auth.Authorization(client, "oleg", "oleg")
	if err != nil {
		log.Printf("authorization error: %s", err)
		return
	}
	fmt.Println("Status code:", response.StatusCode)

}
