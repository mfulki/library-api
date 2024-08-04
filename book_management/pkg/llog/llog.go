package llog

import (
	"book-service/internal/apperror"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	titleInfo  = "[INFO] "
	titleError = "[ERROR] "
	titleFatal = "[FATAL] "
)

var color = map[string]int{
	titleInfo:  32,
	titleError: 31,
	titleFatal: 35,
}

type Logger interface {
	Print(vals ...any)
	Fatal(vals ...any)
}

var logger Logger = NewDefaultLogger()

func SetLogger(newLogger Logger) {
	logger = newLogger
}

func Print(vals ...any) {
	if _, ok := logger.(*defaultLogger); ok {
		vals = append([]any{titleInfo}, vals...)
	}

	logger.Print(vals...)
}

func Error(vals ...any) {
	if _, ok := logger.(*defaultLogger); ok {
		vals = append([]any{titleError}, vals...)
	}

	logger.Print(vals...)
}

func Fatal(vals ...any) {
	if _, ok := logger.(*defaultLogger); ok {
		vals = append([]any{titleFatal}, vals...)
	}

	logger.Fatal(vals...)
}

func timeNow() string {
	return time.Now().Format(time.RFC822)
}

func filePath(vals []any, depthCalled int) string {
	for _, v := range vals[1:] {
		appErr := new(apperror.AppError)
		if err, ok := v.(error); ok && errors.As(err, &appErr) {
			return appErr.FilePath
		}
	}

	if _, file, line, ok := runtime.Caller(depthCalled); ok {
		return fmt.Sprintf("%s:%d", file, line)
	}

	return ""
}

func msg(vals []any) string {
	builder := new(strings.Builder)

	for index, v := range vals {
		builder.WriteString(fmt.Sprint(v))

		if index != len(vals)-1 {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}
