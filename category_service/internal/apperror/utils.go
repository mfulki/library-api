package apperror

import "net/http"

var errStatusCode = map[error]int{
	ErrResourceNotFound: http.StatusBadRequest,
	ErrInvalidRequest:   http.StatusBadRequest,
	ErrNoRoute:          http.StatusNotFound,
	ErrUnauthorized:     http.StatusUnauthorized,
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
