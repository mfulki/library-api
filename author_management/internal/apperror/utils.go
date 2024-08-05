package apperror

import "net/http"

func GetErrStatusCode(sourceErr error) (int, error) {
	var errStatusCode = map[error]int{
		ErrResourceNotFound: http.StatusBadRequest,
		ErrInvalidRequest:   http.StatusBadRequest,
		ErrNoRoute:          http.StatusNotFound,
		ErrUnauthorized:     http.StatusUnauthorized,
	}
	code := http.StatusInternalServerError
	err := ErrInternalServer

	if val, ok := errStatusCode[sourceErr]; ok {
		code = val
		err = sourceErr
	}

	return code, err
}
