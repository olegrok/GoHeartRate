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
	"github.com/olegrok/GoHeartRate/server/workers"
)

var wp *workers.Pool
var requestWaitInQueueTimeout time.Duration

func main() {
	log.Fatalln(runServer())
}

func runServer() error {
	conf := config.LoadConfig("../config.json")
	defer database.Connect().Close()
	wp = workers.NewPool(conf.Options.Concurrency)
	requestWaitInQueueTimeout = conf.Options.RequestWaitInQueueTimeout * time.Second
	wp.Run()

	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	s := &http.Server{
		Addr:           conf.Address,
		Handler:        router,
		ReadTimeout:    conf.Options.ReadTimeout * time.Second,
		WriteTimeout:   conf.Options.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connect! ", *r)

	_, err := wp.AddTaskSyncTimed(func() interface{} {
		var rMsg protocol.ReceivedMessage
		bytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			log.Fatalf("read request body error: %s", err)
		}
		if err := json.Unmarshal(bytes, &rMsg); err != nil {
			log.Fatalf("marshal message error: %s", err)
		}

		switch rMsg.MessageType {
		case "auth":
			var msg protocol.AuthData
			if err := json.Unmarshal(rMsg.Data, &msg); err != nil {
				log.Fatalf("marshal message error: %s", err)
				break
			}
			auth.Authorization(w, msg)
		case "registration":
			var msg protocol.AuthData
			if err := json.Unmarshal(rMsg.Data, &msg); err != nil {
				log.Fatalf("marshal message error: %s", err)
				break
			}
			auth.Registration(w, msg)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Println("unknown message type")
		}
		return nil
	}, requestWaitInQueueTimeout)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}

}
