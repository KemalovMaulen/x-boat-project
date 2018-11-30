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
}

func (e Error) Response() interface{} {
	return e
}
func (e Error) Headers() map[string]string {
	return nil
}

func (e Error) Error() string {
	return fmt.Sprintf("status: %d; code: %s; message: %s; developerMessage: %s", e.Status, e.Code, e.Message, e.DeveloperMessage)
}

func NewErrorSystem(system string) *ErrorSystem {
	return &ErrorSystem{
		System: system,
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

	c := err.System + "."  + strconv.Itoa(code)

	return &Error{
		Status:           status,
		Code:             c,
		Message:          message,
		DeveloperMessage: devMessage,
	}
}

func (errSys *ErrorSystem) BadRequest(code int, messages ...string) *Response {
	resp := &Response{}
	resp.Status = http.StatusBadRequest
	resp.Data = errSys.NewError(resp.Status, code, messages...)
	return resp
}

func (errSys *ErrorSystem) InternalServerError(code int, messages ...string) *Response {
	resp := &Response{}
	resp.Status = http.StatusInternalServerError
	resp.Data = errSys.NewError(resp.Status, code, messages...)
	return resp
}

func (errSys *ErrorSystem) NotFound(code int, messages ...string) *Response {
	resp := &Response{}
	resp.Status = http.StatusNotFound
	resp.Data = errSys.NewError(resp.Status, code, messages...)
	return resp
}

func (errSys *ErrorSystem) Forbidden(code int, messages ...string) *Response {
	resp := &Response{}
	resp.Status = http.StatusForbidden
	resp.Data = errSys.NewError(resp.Status, code, messages...)
	return resp
}

