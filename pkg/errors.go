package pkg

import (
	"fmt"
	"net/http"
)

type ResponseErrors *map[string][]string

type IHttpError interface {
	GetCode() int
	Error() string
	Errors() ResponseErrors
}

type HttpError struct {
	message string
	code    int
	errors  ResponseErrors
}

func NewHttpError(err error, message string, code int, errors ResponseErrors) *HttpError {
	// TODO: add logger
	fmt.Println(err.Error())

	return &HttpError{
		message: message,
		code:    code,
		errors:  errors,
	}
}

func NewCustomHttpError(message string, code int, errors ResponseErrors) *HttpError {
	return &HttpError{
		message: message,
		code:    code,
		errors:  errors,
	}
}

func (e HttpError) Error() string {
	return e.message
}

func (e HttpError) GetCode() int {
	return e.code
}

func (e HttpError) Errors() ResponseErrors {
	return e.errors
}

func NotFoundError(err error, message string) *HttpError {
	return NewHttpError(err, message, http.StatusNotFound, nil)
}

func ConflictError(err error, message string, errors ResponseErrors) *HttpError {
	return NewHttpError(err, message, http.StatusConflict, errors)
}

func InternalServerError(err error, message string) *HttpError {
	return NewHttpError(err, message, http.StatusInternalServerError, nil)
}

func BadRequestError(err error, message string) *HttpError {
	return NewHttpError(err, message, http.StatusBadRequest, nil)
}

// custom errors

func CustomInternalServerError(message string) *HttpError {
	return NewCustomHttpError(message, http.StatusInternalServerError, nil)
}

func CustomUnprocessableEntityError(message string, errors ResponseErrors) *HttpError {
	return NewCustomHttpError(message, http.StatusUnprocessableEntity, errors)
}
