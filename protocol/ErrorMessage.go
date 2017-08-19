package protocol

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"io"
	"errors"
)

const (
	AlreadyRegistered = iota + 1
	WrongPassword
	CalculationError
	JobTimedOut

)

var (
	ErrAlreadyRegistered = errors.New("user has already registered")
	ErrWrongPassword = errors.New("wrong login or password")
	ErrCalculationError = errors.New("calculation server error")
	ErrJobTimedOut = errors.New("request time out")
)

type ErrorData struct {
	Error     error			`json:"error"`
	ErrorCode int			`json:"message_code"`
}

func BytesToErrorData(body io.ReadCloser) ErrorData {
	bytes, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		log.Fatalf("read request body error: %s", err)
	}
	var errorMsg ErrorData
	if err := json.Unmarshal(bytes, &errorMsg); err != nil {
		log.Fatalf("marshal message error: %s", err)
	}
	return errorMsg
}

func ErrorDataToBytes(err error, errCode int) []byte {
	errorMsg := ErrorData {
		Error:     err,
		ErrorCode: errCode,
	}
	data, err := json.Marshal(errorMsg)
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	return []byte(data)
}