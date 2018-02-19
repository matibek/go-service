package core

import (
	"fmt"
)

// ServerError is a warpper over error to allow passing custom values to it.
type ServerError struct {
	Message        string // error message
	Code           int    // app error code
	HttpStatusCode int    // http response code
	Err            error  // the original error
}

func (e *ServerError) Error() (result string) {
	result = fmt.Sprintf("ServerError: %s", e.Message)
	if e.Code != 0 {
		result += fmt.Sprintf(" [code=%d]\n", e.Code)
	}
	if e.Err != nil {
		result += e.Err.Error()
	}
	return
}

// NewError makes a ServerError instance from the given value
func NewError(msg string, err error) *ServerError {
	return &ServerError{
		Message:        msg,
		Err:            err,
		HttpStatusCode: 500,
	}
}
