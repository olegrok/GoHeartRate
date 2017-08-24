package protocol

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

// Codes of errors
const (
	AlreadyRegistered = iota + 1
	WrongPassword
	CalculationError
	JobTimedOut
	DatabaseError
)

// Errors descriptions
const (
	ErrAlreadyRegistered = "user has already registered"
	ErrWrongPassword     = "wrong login or password"
	ErrCalculation       = "calculation server error"
	ErrJobTimedOut       = "request time out"
	ErrDatabase          = "database error"
)

// ErrorData contains errors code and description
type ErrorData struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

// BytesToErrorData converts body of response to ErrorData structure
func BytesToErrorData(body io.ReadCloser) ErrorData {
	bytes, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		log.Fatalf("read request body error: %s", err)
	}
	var errorMsg ErrorData
	if err := json.Unmarshal(bytes, &errorMsg); err != nil {
		log.Fatalf("unmarshal message error: %s", err)
	}
	return errorMsg
}

// ErrorDataToBytes converts ErrorData structure to bytes
func ErrorDataToBytes(er string, errCode int) []byte {
	data, err := json.Marshal(ErrorData{
		Error:     er,
		ErrorCode: errCode,
	})
	if err != nil {
		log.Printf("marshaling error: %s", err)
	}
	return data
}
