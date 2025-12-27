package types

import "net/http"

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *HTTPError {
	return &HTTPError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewNotFoundError(message string) *HTTPError {
	return &HTTPError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *HTTPError {
	return &HTTPError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewServiceUnavailableError(message string) *HTTPError {
	return &HTTPError{
		Code:    http.StatusServiceUnavailable,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *HTTPError {
	return &HTTPError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewUnprocessableEntityError(message string) *HTTPError {
	return &HTTPError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
