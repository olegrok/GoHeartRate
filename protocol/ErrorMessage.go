package protocol

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

const (
	AlreadyRegistered = iota + 1
	WrongPassword
	CalculationError
	JobTimedOut
	DatabaseError
)

var (
	ErrAlreadyRegistered = "user has already registered"
	ErrWrongPassword     = "wrong login or password"
	ErrCalculation       = "calculation server error"
	ErrJobTimedOut       = "request time out"
	ErrDatabase          = "database error"
)

type ErrorData struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

func BytesToErrorData(body io.ReadCloser) ErrorData {
	bytes, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		log.Fatalf("read request body error: %s", err)
	}
	var errorMsg ErrorData
	fmt.Printf("%s\n", bytes)
	if err := json.Unmarshal(bytes, &errorMsg); err != nil {
		log.Fatalf("unmarshal message error: %s", err)
	}
	return errorMsg
}

func ErrorDataToBytes(er string, errCode int) []byte {
	data, err := json.Marshal(ErrorData{
		Error:     er,
		ErrorCode: errCode,
	})
	fmt.Printf("%s %s\n", data, er)
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	return data
}
