package custom_errors

import (
	"errors"
	"fmt"
)

//Errors detected: BadRequest, NotFoundError, InternalError, UnauthorizedError

type ErrorType string

const (
	ErrBadRequestType          ErrorType = "BAD_REQUEST"
	ErrNotFoundType            ErrorType = "NOT_FOUND"
	ErrUnauthorizedType        ErrorType = "UNAUTHORIZED"
	ErrSaveType                ErrorType = "SAVE_ERROR"
	ErrInternalServerErrorType ErrorType = "INTERNAL_SERVER_ERROR"
	ErrDuplicatedKeyType       ErrorType = "DUPLICATE_KEY"
)

var ErrBadRequest = errors.New("the request is invalid or malformed")
var ErrNotFound = errors.New("error data not found")
var ErrUnauthorized = errors.New("the user is not authorized")
var ErrSavingError = errors.New("error saving refresh token")
var ErrInternalServerError = errors.New("internal server error")
var ErrDuplicatedKey = errors.New("duplicate key")

type LogError struct {
	Type       ErrorType
	UUID       string
	Message    string
	StatusCode int
	LogMessage string
}

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func New(errType ErrorType, uuid string, message string, statusCode int, logMessage string) *LogError {
	return &LogError{
		Type:       errType,
		UUID:       uuid,
		Message:    message,
		StatusCode: statusCode,
		LogMessage: logMessage,
	}
}

func (e *LogError) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}
