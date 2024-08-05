package validate

import (
	"author-service/internal/apperror"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func init() {
	validate.RegisterTagNameFunc(fieldTagNew)
	validate.RegisterValidation("date", isDateTimeFormat(time.DateOnly))
	validate.RegisterValidation("timestamp", isDateTimeFormat(time.RFC3339))
}

func BodyJSON(ctx *fiber.Ctx, reqBody any) error {
	ctx.Context().Request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return bodyParser(ctx, reqBody)
}

func BodyFormData(ctx *fiber.Ctx, reqBody any) error {
	ctx.Context().Request.Header.Set(fiber.HeaderContentType, fiber.MIMEMultipartForm)
	return bodyParser(ctx, reqBody)
}

func bodyParser(ctx *fiber.Ctx, reqBody any) error {
	if err := ctx.BodyParser(reqBody); err != nil {
		return apperror.Wrap(err)
	}

	if err := validate.Struct(reqBody); err != nil {
		return apperror.Wrap(err)
	}

	return nil
}

func fieldTag(field reflect.StructField, tagName string) string {
	name := strings.SplitN(field.Tag.Get(tagName), ",", 2)[0]
	if name == "-" {
		return ""
	}

	return name
}

func fieldTagNew(field reflect.StructField) string {
	name := fieldTag(field, "json")
	if name == "" {
		name = fieldTag(field, "form")
	}

	if name == "" {
		name = fieldTag(field, "uri")
	}

	return name
}

func isDateTimeFormat(format string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		_, err := time.Parse(format, value)

		return err == nil
	}
}
