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

type StatusError struct {
	err     error
	message string
	status  int
}

func (e StatusError) Error() string {
	return fmt.Sprintf("[%d] %s", e.status, e.message)
}

func (e StatusError) Is(status int) bool {
	return e.status == status
}

func (e StatusError) In(statuses ...int) bool {
	for _, status := range statuses {
		if e.Is(status) {
			return true
		}
	}
	return false
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
