package apperror

import "net/http"

var errStatusCode = map[error]int{
	ErrResourceNotFound:  http.StatusBadRequest,
	ErrInvalidRequest:    http.StatusBadRequest,
	ErrNoRoute:           http.StatusNotFound,
	ErrUnauthorized:      http.StatusUnauthorized,
	ErrInvalidCredential: http.StatusBadRequest,
	ErrInvalidToken:      http.StatusBadRequest,
	ErrInvalidEmail:      http.StatusBadRequest,
	ErrInvalidPassword:   http.StatusBadRequest,
	ErrInvalidPassToken:  http.StatusBadRequest,
	ErrInvalidParam:      http.StatusBadRequest,
}

func GetErrStatusCode(sourceErr error) (int, error) {
	code := http.StatusInternalServerError
	err := ErrInternalServer

	if val, ok := errStatusCode[sourceErr]; ok {
		code = val
		err = sourceErr
	}

	return code, err
}
