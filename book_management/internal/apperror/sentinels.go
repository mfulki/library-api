package apperror

import (
	"errors"
)

var (
	ErrNoRoute          = errors.New("route is not found")
	ErrInternalServer   = errors.New("the server encounter error, please try again later")
	ErrResourceNotFound = errors.New("resource is not found")
	ErrInvalidRequest   = errors.New("payload is invalid")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrAssertingAny     = errors.New("cannot asserting")
	ErrCannotBorrowUnAvailable=errors.New("cannot borrow an unavailable book")
)
