package api

import (
	"fmt"
	"net/http"
)

var (
	ErrBadRequest          = StatusError{status: http.StatusBadRequest, message: "bad request"}
	ErrInternalServerError = StatusError{status: http.StatusInternalServerError, message: "internal server error"}
	ErrMethodNotAllowed    = StatusError{status: http.StatusMethodNotAllowed, message: "method not allowed"}
)

// StatusError
type StatusError struct {
	err     error  // for internal logging
	message string // show to user
	status  int    // HTTP status code
}

func (e StatusError) Error() string {
	return fmt.Sprintf("[%d] %s", e.status, e.message)
}

func (e StatusError) Status() int {
	return e.status
}

func (e StatusError) With(err error) StatusError {
	e.err = err
	return e
}

func ToStatusError(err error) StatusError {
	switch t := err.(type) {
	case StatusError:
		return t
	case error:
		return ErrInternalServerError.With(t)
	default:
		panic("unknown type")
	}
}
