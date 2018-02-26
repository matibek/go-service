package core

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	libErrors "github.com/pkg/errors"
)

// Error is a warpper over error to allow passing custom values to it.
type Error struct {
	Message        string // error message
	Code           int    // app error code
	HttpStatusCode int    // http response code
	Err            error  // the original error
}

// New creates an error with stack trace and http response code
func (e *Error) New(err error) *Error {
	if err == nil {
		return nil
	}
	return e.wrap(err, "")
}

// NewWithMessage creates error with custom message for response
func (e *Error) NewWithMessage(err error, message string) *Error {
	if err == nil {
		return nil
	}
	return e.wrap(err, message)
}

// NewWithMessagef creates error with formatted custom message for response
func (e *Error) NewWithMessagef(err error, format string, args ...interface{}) *Error {
	if err == nil {
		return nil
	}
	return e.wrap(err, fmt.Sprintf(format, args...))
}

// NewError creates a new error from message
func (e *Error) NewError(message string) *Error {
	err := errors.New(message) // go error
	return e.wrap(err, message)
}

// NewErrorf creates a new error from fromatted message
func (e *Error) NewErrorf(format string, args ...interface{}) *Error {
	message := fmt.Sprintf(format, args...)
	err := libErrors.Errorf(format, args...)
	return e.wrap(err, message)
}

func (*Error) wrap(err error, message string) *Error {
	err = libErrors.WithStack(err) // apply stack trace
	return &Error{
		Message:        message,
		Err:            err,
		HttpStatusCode: http.StatusInternalServerError,
	}
}

func (e *Error) Error() (result string) {
	result = fmt.Sprintf("Error: %s", e.Message)
	if e.Code != 0 {
		result += fmt.Sprintf(" [code=%d]", e.Code)
	}
	result += "\n" + e.Err.Error()
	return
}

func (e *Error) Format(s fmt.State, verb rune) {
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

func (e *Error) JSON(stack bool) interface{} {
	json := gin.H{}
	json["error"] = e.Message
	json["code"] = e.Code
	if stack {
		json["stack"] = fmt.Sprintf("%+v", e)
	}
	return json
}
