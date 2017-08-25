package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/olegrok/GoHeartRate/protocol"
	"github.com/olegrok/GoHeartRate/server/auth"
	"github.com/olegrok/GoHeartRate/server/config"
	"github.com/olegrok/GoHeartRate/server/database"
	"github.com/olegrok/GoHeartRate/server/math"
	"github.com/olegrok/GoHeartRate/server/workers"
)

var wp *workers.Pool
var requestWaitInQueueTimeout time.Duration

func main() {
	defer database.DB.Close()
	log.Fatalln(runServer())
}

func init() {
	config.LoadConfig("../config.json")
	database.Connect()
}

func runServer() error {
	wp = workers.NewPool(config.Config.Options.Concurrency)
	requestWaitInQueueTimeout = config.Config.Options.RequestWaitInQueueTimeout * time.Second
	wp.Run()

	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	s := &http.Server{
		Addr:           config.Config.Address,
		Handler:        router,
		ReadTimeout:    config.Config.Options.ReadTimeout * time.Second,
		WriteTimeout:   config.Config.Options.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}

func messageDistributor(rMsg protocol.ReceivedMessage, w http.ResponseWriter, r *http.Request) {
	switch rMsg.MessageType {
	case protocol.Auth:
		if msg, err := auth.GetAuthMessage(rMsg.Data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			auth.Authorization(w, *msg)
		}
	case protocol.Register:
		if msg, err := auth.GetAuthMessage(rMsg.Data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			auth.Registration(w, *msg)
		}
	case protocol.Data:
		if uid, status := database.IsAuthorizedUser(r.Cookies()); status {
			if err := math.ResultHandler(rMsg.Data, uid, w); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				break
			}
			w.WriteHeader(http.StatusOK)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	case protocol.Results:
		if uid, status := database.IsAuthorizedUser(r.Cookies()); status {
			if res, err := database.GetResults(uid); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				break
			} else {
				data, err := json.Marshal(res)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					break
				}
				w.Write(data)
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}

	default:
		w.WriteHeader(http.StatusNotAcceptable)
		log.Printf("unknown message type: %s", rMsg.MessageType)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connect!", *r)
	_, err := wp.AddTaskSyncTimed(func() interface{} {
		var rMsg protocol.ReceivedMessage
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("read request body error: %s", err)
		}
		defer r.Body.Close()

		if err = json.Unmarshal(bytes, &rMsg); err != nil {
			log.Printf("marshal message error: %s", err)
		}
		messageDistributor(rMsg, w, r)
		return nil
	}, requestWaitInQueueTimeout)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}

}
