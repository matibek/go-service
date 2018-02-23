package core

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
	result += "\n" + e.Err.Error()
	return
}

func (e *ServerError) Format(s fmt.State, verb rune) {
	io.WriteString(s, e.Message+"\n")
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", e.Err)
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(s, "%s", e.Err)
	case 'q':
		fmt.Fprintf(s, "%q", e.Err)
	}
}

func (e *ServerError) JSON(stack bool) interface{} {
	json := gin.H{}
	json["error"] = e.Message
	json["code"] = e.Code
	if stack {
		json["stack"] = fmt.Sprintf("%+v", e)
	}
	return json
}

// NewError makes a ServerError instance from the given value
func NewError(message string, err error) *ServerError {
	if err == nil {
		err = errors.New(message)
	} else {
		err = errors.WithStack(err)
	}
	return &ServerError{
		Message:        message,
		Err:            err,
		HttpStatusCode: http.StatusInternalServerError,
	}
}
