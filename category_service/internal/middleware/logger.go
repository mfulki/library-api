package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"category-service/internal/apperror"
	"category-service/internal/constant"
	"category-service/pkg/llog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Fields struct {
	Type       string `json:"type"`
	RequestID  string `json:"requestId"`
	Latency    string `json:"latency"`
	Method     string `json:"method"`
	PathURL    string `json:"url"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	File       string `json:"file,omitempty"`
}

func (m *Middleware) loggerError(ctx *fiber.Ctx, err error) {
	statusCode := ctx.Response().StatusCode()

	if !(statusCode >= 500 && statusCode <= 599) {
		return
	}

	uri := ctx.Request().URI()
	path := string(uri.Path())
	rawQuery := string(uri.QueryString())

	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}

	start := ctx.Locals(startTime).(time.Time)
	idAny := ctx.Locals(constant.RequestIDKey)
	id, _ := idAny.(string)

	flds := Fields{
		Type:       "ERROR",
		RequestID:  id,
		Latency:    fmt.Sprint(time.Since(start)),
		Method:     string(ctx.Request().Header.Method()),
		PathURL:    path,
		StatusCode: statusCode,
		Message:    fmt.Sprint(err),
	}

	if appErr := new(apperror.AppError); errors.As(err, &appErr) {
		flds.File = appErr.FilePath
	}

	llog.Error(err)
	jsonByte, _ := json.Marshal(flds)
	m.fileLogger.Print(string(jsonByte))
}

func (m *Middleware) LoggerInfo(ctx *fiber.Ctx) error {
	nextErr := ctx.Next()
	statusCode := ctx.Response().StatusCode()

	if nextErr != nil {
		code, err := apperror.GetErrStatusCode(nextErr)
		statusCode = code

		if ve := make(validator.ValidationErrors, 0); errors.Is(err, apperror.ErrInternalServer) {
			if !errors.As(nextErr, &ve) {
				return nextErr
			}

			statusCode = http.StatusBadRequest
		}

	}

	uri := ctx.Request().URI()
	path := string(uri.Path())
	rawQuery := string(uri.QueryString())

	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}

	start := ctx.Locals(startTime).(time.Time)
	idAny := ctx.Locals(constant.RequestIDKey)
	id, _ := idAny.(string)

	flds := Fields{
		Type:       "INFO",
		RequestID:  id,
		Latency:    fmt.Sprint(time.Since(start)),
		Method:     string(ctx.Request().Header.Method()),
		PathURL:    path,
		StatusCode: statusCode,
		Message:    "request processed",
	}

	jsonByte, _ := json.Marshal(flds)
	m.fileLogger.Print(string(jsonByte))

	return nextErr
}
