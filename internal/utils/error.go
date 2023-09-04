package utils

import "errors"

var (
	ErrInvalidInputFormat  = errors.New("invalid request body or query parameters")
	ErrInternalServerError = errors.New("internal server error")
	ErrStudentNotFound     = errors.New("student not found")
	ErrTeacherNotFound     = errors.New("teacher not found")
	ErrDuplicateEntry      = errors.New("email already taken")
)
