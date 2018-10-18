package errors

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
