package utils

import "errors"

var (
	ErrInvalidInputFormat  = errors.New("invalid request body or query parameters")
	ErrInternalServerError = errors.New("internal server error")
)
