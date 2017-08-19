package main

import (
	"fmt"
	"github.com/olegrok/GoHeartRate/client/auth"
	"log"
	"net/http"
	"time"
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
