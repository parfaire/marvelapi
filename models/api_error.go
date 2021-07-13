package models

import (
	"net/http"

	"github.com/parfaire/marvelapi/util"
)

type ApiError struct {
	Message    string
	HTTPStatus int
}

func (apiError *ApiError) Error() string {
	return apiError.Message
}

// generic use of error 500 with custom message as param
func NewApiError(message string) *ApiError {
	return &ApiError{
		Message:    message,
		HTTPStatus: http.StatusInternalServerError,
	}
}

// function convetion : to always have Error as prefix
// eg. ErrorInternalServer, ErrorMissingRequiredParameter, ErrorXXXXXXX, and so on..
func ErrorNotFound() *ApiError {
	return &ApiError{
		Message:    util.ERROR_NOT_FOUND,
		HTTPStatus: http.StatusNotFound,
	}
}

func ErrorInternalServer() *ApiError {
	return &ApiError{
		Message:    util.ERROR_INTERNAL_SERVER,
		HTTPStatus: http.StatusInternalServerError,
	}
}
