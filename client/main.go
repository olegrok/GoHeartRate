package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/olegrok/GoHeartRate/client/auth"
	"github.com/olegrok/GoHeartRate/client/math"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("client error: %s", err)
	}

	client := &http.Client{
		Transport: tr,
		Jar:       jar,
	}

	//auth.Registration(client, "oleg3", "pass")

	res, err := auth.StartLogin(client)
	if err != nil {
		log.Printf("authorization error: %s", err)
		return
	}
	fmt.Println("Login status code:", res.StatusCode)
	fmt.Println(*res, res.Cookies())

	res, err = math.Transmit(client, []float64{1, 2.71, 3.14})

	fmt.Println("Transmit results status code:", res.StatusCode)
	fmt.Println(*res, res.Cookies())

	bytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Printf("%s\n", bytes)
	//response, err := auth.Authorization(client, "oleg", "oleg")

}
