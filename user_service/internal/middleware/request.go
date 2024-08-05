package middleware

import (
	"fmt"
	"time"
	"user-service/internal/constant"

	"github.com/gofiber/fiber/v2"
)

const startTime = "startTime"

func (m *Middleware) RequestID(ctx *fiber.Ctx) error {
	timeNow := time.Now()
	reqID := fmt.Sprintf("%s-%d", constant.RequestIDPrefix, timeNow.Unix())

	ctx.Set(constant.RequestIDKey, reqID)

	ctx.Locals(constant.RequestIDKey, reqID)
	ctx.Locals(startTime, timeNow)

	return ctx.Next()
}
