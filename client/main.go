package main

import (
	"encoding/json"
	"fmt"
	"github.com/olegrok/GoHeartRate/client/auth"
	"github.com/olegrok/GoHeartRate/client/math"
	"github.com/olegrok/GoHeartRate/client/requests"
	"github.com/olegrok/GoHeartRate/server/database"
	"github.com/olegrok/GoHeartRate/webcam"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
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

	res := auth.StartLogin(client)
	fmt.Println("Login status code:", res.StatusCode)
	signal, time_array := webcam.Start()
	if res, err = math.Transmit(client, signal, time_array); err != nil {
		//if res, err = math.Transmit(client, []float64{1, 2.71, 3.14}); err != nil {
		log.Fatalf("result transmitting error: %s", err)
	}
	fmt.Println("Transmit results status code:", res.StatusCode)
	if bytes, err := ioutil.ReadAll(res.Body); err != nil {
		log.Fatalf("response reading error: %s", err)
	} else {
		fmt.Printf("YOUR RESULT: %s\n", bytes)
	}
	defer res.Body.Close()

	if res, err = requests.GetLastResults(client); err != nil {
		log.Fatalf("reading results error: %s", err)
	}

	if bytes, err := ioutil.ReadAll(res.Body); err != nil {
		log.Fatalf("response reading error: %s", err)
	} else {
		fmt.Println("YOUR LAST 10 RESULTS:")
		var results []database.UserResult
		json.Unmarshal(bytes, &results)
		for _, result := range results {
			fmt.Printf("%s: %s\n", result.CreatedAt.Format("02.01.2006 15:04"), result.Result)
		}
	}

}
