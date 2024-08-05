package apperror

import (
	"fmt"
	"runtime"
)

type AppError struct {
	FilePath string
	Err      error
}

func Wrap(err error) *AppError {
	filePath := ""
	if _, file, line, ok := runtime.Caller(1); ok {
		filePath = fmt.Sprintf("%s:%d", file, line)
	}

	return &AppError{
		FilePath: filePath,
		Err:      err,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}
