package server

import (
	"fmt"
	"net/http"
	"strconv"
)

type Error struct {
	Status           int    `json:"status,omitempty"`
	Code             string `json:"code,omitempty"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
}

type ErrorSystem struct {
	System string `json:"system,omitempty"`
	Series int    `json:"series,omitempty"`
}

type (
	DarError interface {
		Headers() map[string]string

		Error() string

		Response() interface{}

		StatusCode() int
	}
)

func (e Error) Response() interface{} {
	return e
}
func (e Error) Headers() map[string]string {
	return nil
}

func (e Error) Error() string {
	return fmt.Sprintf("status: %d; code: %s; message: %s; developerMessage: %s", e.Status, e.Code, e.Message, e.DeveloperMessage)
}

func NewErrorSystem(system string, series int) *ErrorSystem {
	return &ErrorSystem{
		System: system,
		Series: series,
	}
}

func (err *ErrorSystem) NewError(status, code int, messages ...string) *Error {
	var message, devMessage string

	switch len(messages) {
	case 2:
		message, devMessage = messages[0], messages[1]
	case 1:
		message = messages[0]
	}

	c := strconv.Itoa(status)+ "." + strconv.Itoa(code)

	return &Error{
		Status:           status,
		Code:             c,
		Message:          message,
		DeveloperMessage: devMessage,
	}
}

func (errSys *ErrorSystem) BadRequest(code int, messages ...string) *Error {
	return errSys.NewError(http.StatusBadRequest, code, messages...)
}

func (errSys *ErrorSystem) InternalServerError(code int, messages ...string) *Error {
	return errSys.NewError(http.StatusInternalServerError, code, messages...)
}

func (errSys *ErrorSystem) NotFound(code int, messages ...string) *Error {
	return errSys.NewError(http.StatusNotFound, code, messages...)
}

func (errSys *ErrorSystem) Forbidden(code int, messages ...string) *Error {
	return errSys.NewError(http.StatusForbidden, code, messages...)
}

