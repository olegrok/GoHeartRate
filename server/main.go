package main

import (
	"fmt"
	//"encoding/json"
	"log"
	"github.com/olegrok/GoHeartRate/protocol"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/olegrok/GoHeartRate/server/auth"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now(), ": Connect! ", r)
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
		auth.Authorization(w, r, msg)

	default:
		fmt.Println("unknown message type")
	}
	fmt.Fprint(w, "")
}
