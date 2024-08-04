package apperror

import (
	"github.com/gofiber/fiber/v2"
)

var errStatusCode = map[error]int{
	ErrResourceNotFound:  fiber.StatusBadRequest,
	ErrInvalidRequest:    fiber.StatusBadRequest,
	ErrNoRoute:           fiber.StatusNotFound,
	ErrUnauthorized:      fiber.StatusUnauthorized,
	ErrInvalidCredential: fiber.StatusBadRequest,
	ErrInvalidToken:      fiber.StatusBadRequest,
	ErrInvalidEmail:      fiber.StatusBadRequest,
	ErrInvalidPassword:   fiber.StatusBadRequest,
	ErrInvalidPassToken:  fiber.StatusBadRequest,
	ErrInvalidParam:      fiber.StatusBadRequest,
	ErrEmailExist:        fiber.StatusBadRequest,
}

func GetErrStatusCode(sourceErr error) (int, error) {
	code := fiber.StatusInternalServerError
	err := ErrInternalServer

	if val, ok := errStatusCode[sourceErr]; ok {
		code = val
		err = sourceErr
	}

	return code, err
}
