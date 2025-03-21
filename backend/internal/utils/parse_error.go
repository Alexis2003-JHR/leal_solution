package utils

import (
	"errors"
	"leal/internal/core/domain/custom_errors"
	"net/http"
)

func ParseError(err error) *custom_errors.LogError {
	var customErr *custom_errors.LogError
	switch {
	case errors.Is(err, custom_errors.ErrBadRequest):
		customErr = custom_errors.New(custom_errors.ErrBadRequestType, "Invalid request", http.StatusBadRequest, err.Error())

	case errors.Is(err, custom_errors.ErrDuplicatedKey):
		customErr = custom_errors.New(custom_errors.ErrDuplicatedKeyType, "The entity already exists", http.StatusConflict, err.Error())

	case errors.Is(err, custom_errors.ErrUnauthorized):
		customErr = custom_errors.New(custom_errors.ErrUnauthorizedType, "Unauthorized", http.StatusUnauthorized, err.Error())

	case errors.Is(err, custom_errors.ErrNotFound):
		customErr = custom_errors.New(custom_errors.ErrNotFoundType, "Entity not found", http.StatusNotFound, err.Error())

	default:
		customErr = custom_errors.New(custom_errors.ErrInternalServerErrorType, "Something unexpected has happened", http.StatusInternalServerError, err.Error())
	}

	return customErr
}
