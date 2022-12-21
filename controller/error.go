package controller

// ErrorCode Use type alias to simulate enum
type ErrorCode string

const (
	ErrMissingAuth  = "MISSING_AUTH"
	ErrBadAuth      = "BAD_AUTH"
	ErrJwtSignError = "SIGN_JWT_ERROR"
	ErrServerError  = "SERVER_ERROR"
)

// HttpError The standard error format
type HttpError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}
