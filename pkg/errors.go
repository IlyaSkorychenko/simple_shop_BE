package pkg

import (
	"fmt"
	"net/http"
)

type IHttpError interface {
	GetCode() int
	Error() string
}

type HttpError struct {
	message string
	code    int
}

func NewHttpError(err error, message string, code int) *HttpError {
	// TODO: add logger
	fmt.Println(err.Error())

	return &HttpError{
		message: message,
		code:    code,
	}
}

func (e HttpError) Error() string {
	return e.message
}

func (e HttpError) GetCode() int {
	return e.code
}

func NewNotFoundError(err error, message string) *HttpError {
	return NewHttpError(err, message, http.StatusNotFound)
}

func NewInternalServerError(err error, message string) *HttpError {
	return NewHttpError(err, message, http.StatusInternalServerError)
}
