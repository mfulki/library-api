package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"user-service/internal/apperror"
	"user-service/internal/dto"
	"user-service/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) ErrorHandler(ctx *fiber.Ctx, err error) error {
	defer m.loggerError(ctx, err)
	if ve := make(validator.ValidationErrors, 0); errors.As(err, &ve) {
		mapErrors := make(map[string]string)
		for _, fe := range ve {
			mapErrors[fe.Field()] = utils.GetValidationErrorMsg(fe)
		}

		return ctx.Status(http.StatusBadRequest).JSON(dto.Response{
			Message: fmt.Sprint(apperror.ErrInvalidRequest),
			Errors:  mapErrors,
		})
	}

	statusCode, err := apperror.GetErrStatusCode(err)

	return ctx.Status(statusCode).JSON(dto.Response{
		Message: err.Error(),
	})
}
