package main

import (
	"fmt"
	"log"
	"github.com/olegrok/GoHeartRate/protocol"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/olegrok/GoHeartRate/server/auth"
	"time"
	"github.com/olegrok/GoHeartRate/server/workers"
)

const requestWaitInQueueTimeout = time.Second * 15
const kernels = 8
var wp = workers.NewPool(kernels)

func main() {

	wp.Run()
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	s := &http.Server {
		Addr: protocol.Addr,
		Handler: router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
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
