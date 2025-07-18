package rest_err

import "net/http"

type Errors struct {
	Message string    `json:"message"`
	Err     string    `json:"error"`
	Code    int       `json:"code"`
	Details []Details `json:"details,omitempty"`
}

type Details struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *Errors) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *Errors {
	return &Errors{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedError(message string) *Errors {
	return &Errors{
		Message: message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewValidationError(message string, details []Details) *Errors {
	return &Errors{
		Message: message,
		Err:     "Validation Error",
		Code:    http.StatusBadRequest,
		Details: details,
	}
}

func NewInternalServerError(message string) *Errors {
	return &Errors{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Errors {
	return &Errors{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *Errors {
	return &Errors{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}
