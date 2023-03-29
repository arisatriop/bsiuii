package bsiuii

import "net/http"

type Error struct {
	Code    int
	Message interface{}
	Data    interface{}
}

func ServerError() *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: "terjadi kesalahan pada server",
	}
}

func BadRequest(m string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: m,
	}
}

func NotFound(m ...interface{}) *Error {
	var message interface{}
	message = "The data not found."
	if len(m) > 0 {
		message = m[0]
	}
	return &Error{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
