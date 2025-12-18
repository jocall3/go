package errors

import (
	"fmt"
	"net/http"
)

// Error represents a standardized application error that can be serialized to JSON.
// It implements the standard error interface.
type Error struct {
	// Status is the HTTP status code associated with the error.
	Status int `json:"status"`

	// Message is a human-readable description of the error.
	Message string `json:"message"`

	// Details provides optional additional context (e.g., validation errors).
	Details interface{} `json:"details,omitempty"`

	// Cause is the underlying error that triggered this one, if any.
	// It is excluded from JSON serialization to prevent leaking internal details.
	Cause error `json:"-"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap returns the underlying cause of the error, allowing usage with errors.Is and errors.As.
func (e *Error) Unwrap() error {
	return e.Cause
}

// New creates a new Error with the given status code and message.
func New(status int, message string, args ...interface{}) *Error {
	return &Error{
		Status:  status,
		Message: fmt.Sprintf(message, args...),
	}
}

// Wrap creates a new Error wrapping an existing error with a status code and message.
func Wrap(err error, status int, message string, args ...interface{}) *Error {
	return &Error{
		Status:  status,
		Message: fmt.Sprintf(message, args...),
		Cause:   err,
	}
}

// WithDetails adds additional details to the error (e.g., validation map) and returns the error.
func (e *Error) WithDetails(details interface{}) *Error {
	e.Details = details
	return e
}

// NewBadRequest creates a 400 Bad Request error.
func NewBadRequest(message string, args ...interface{}) *Error {
	return New(http.StatusBadRequest, message, args...)
}

// NewUnauthorized creates a 401 Unauthorized error.
func NewUnauthorized(message string, args ...interface{}) *Error {
	return New(http.StatusUnauthorized, message, args...)
}

// NewForbidden creates a 403 Forbidden error.
func NewForbidden(message string, args ...interface{}) *Error {
	return New(http.StatusForbidden, message, args...)
}

// NewNotFound creates a 404 Not Found error.
func NewNotFound(message string, args ...interface{}) *Error {
	return New(http.StatusNotFound, message, args...)
}

// NewConflict creates a 409 Conflict error.
func NewConflict(message string, args ...interface{}) *Error {
	return New(http.StatusConflict, message, args...)
}

// NewInternalServerError creates a 500 Internal Server Error.
// It automatically wraps the original error.
func NewInternalServerError(err error, message string, args ...interface{}) *Error {
	if message == "" {
		message = "internal server error"
	}
	return Wrap(err, http.StatusInternalServerError, message, args...)
}

// GetStatus returns the HTTP status code for a given error.
// If the error is of type *Error, it returns its Status.
// Otherwise, it defaults to 500 Internal Server Error.
func GetStatus(err error) int {
	if e, ok := err.(*Error); ok {
		return e.Status
	}
	return http.StatusInternalServerError
}