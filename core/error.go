package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
		result += fmt.Sprintf(" [code=%d]", e.Code)
	}
	if e.Err != nil {
		result += "\n" + e.Err.Error()
	}
	return
}

func (e *ServerError) JSON() interface{} {
	json := gin.H{}
	json["error"] = e.Message
	json["detail"] = e.Error() // TODO show only incase of debug
	return json
}

// NewError makes a ServerError instance from the given value
func NewError(msg string, err error) *ServerError {
	return &ServerError{
		Message:        msg,
		Err:            err,
		HttpStatusCode: http.StatusInternalServerError,
	}
}
