package errors

import (
	"fmt"
	"net/http"
)

type RestError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restError struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restError) Message() string {
	return e.ErrMessage
}

func (e restError) Status() int {
	return e.ErrStatus
}

func (e restError) Causes() []interface{} {
	return e.ErrCauses
}

func NewBadRequestError(message string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewInternalServerError(message string, err error) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
}
