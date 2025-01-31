package errors

import "errors"

var (
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")
	ErrDBFail       = errors.New("database operation failed")
	ErrDBInit       = errors.New("database init failed")
)
