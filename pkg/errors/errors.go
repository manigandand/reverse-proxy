package errors

import "net/http"

// AppError struct holds the value of HTTP status code and custom error message.
type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"error_message,omitempty"`
}

func (err *AppError) Error() string {
	return err.Message
}

// NewAppError returns the new apperror object
func NewAppError(status int, message string) *AppError {
	return &AppError{status, message}
}

// 4xx -------------------------------------------------------------------------

// BadRequest will return `http.StatusBadRequest` with custom message.
func BadRequest(message string) *AppError { // 400
	return &AppError{http.StatusBadRequest, message}
}

// NotFound will return `http.StatusNotFound` with custom message.
func NotFound(message string) *AppError { // 404
	return &AppError{http.StatusNotFound, message}
}

// UnprocessableEntity will return `http.StatusUnprocessableEntity` with
// custom message.
func UnprocessableEntity(message string) *AppError { // 422
	return &AppError{http.StatusUnprocessableEntity, message}
}

// 5xx -------------------------------------------------------------------------

// InternalServer will return `http.StatusInternalServerError` with custom message.
func InternalServer(message string) *AppError { // 500
	return &AppError{http.StatusInternalServerError, message}
}

// IsStatusNotFound should return true if HTTP status of an error is 404.
func (err *AppError) IsStatusNotFound() bool {
	return err.Status == http.StatusNotFound
}
