package apperror

import (
	"errors"
)

var (
	ErrNoRoute           = errors.New("route is not found")
	ErrInternalServer    = errors.New("the server encounter error, please try again later")
	ErrResourceNotFound  = errors.New("resource is not found")
	ErrInvalidRequest    = errors.New("payload is invalid")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidCredential = errors.New("email or password is incorrect")
	ErrInvalidEmail      = errors.New("email is invalid")
	ErrInvalidToken      = errors.New("token is invalid")
	ErrInvalidPassword   = errors.New("password is incorrect")
	ErrInvalidPassToken  = errors.New("password or token combination is incorrect")
	ErrInvalidParam      = errors.New("route param is invalid")
	ErrEmailExist        = errors.New("the email already registered")
	ErrAssertingAny      = errors.New("cannot asserting")
)
