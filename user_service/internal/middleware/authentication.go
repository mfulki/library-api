package middleware

import (
	"strings"
	"user-service/internal/apperror"
	"user-service/internal/constant"
	"user-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (m *Middleware) Authentication(jwtFunc func(signed string) (jwt.MapClaims, bool)) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")
		bearerToken := strings.Split(authorization, " ")

		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			return apperror.ErrUnauthorized
		}

		user, ok := jwtFunc(bearerToken[1])
		if !ok {
			return apperror.ErrUnauthorized
		}

		userDataMap, ok := user["data"].(map[string]any)
		if !ok {
			return apperror.ErrUnauthorized
		}

		if role, ok := user["role"]; ok {
			userDataMap["role"] = role
		}

		ctx.Locals(constant.UserContext, userDataMap)
		return ctx.Next()
	}
}
func (m *Middleware) UserAuth() fiber.Handler {
	return m.Authentication(func(signed string) (jwt.MapClaims, bool) {
		user, ok := utils.JwtParseUser(signed)
		if !ok {
			return nil, false
		}

		user["role"] = "user"
		return user, true
	})
}

func (m *Middleware) AdminAuth() fiber.Handler {
	return m.Authentication(func(signed string) (jwt.MapClaims, bool) {
		admin, ok := utils.JwtParseAdmin(signed)
		if !ok {
			return nil, false
		}

		admin["role"] = "admin"
		return admin, true
	})
}
func (m *Middleware) AuthMulti(actors []string) fiber.Handler {
	return m.Authentication(func(signed string) (jwt.MapClaims, bool) {
		for _, actor := range actors {
			if act, ok := m.isActorAuthed(actor, signed); ok {
				act["role"] = actor
				return act, ok
			}
		}

		return nil, false
	})
}

func (m *Middleware) isActorAuthed(actor string, signed string) (jwt.MapClaims, bool) {
	act, ok := make(jwt.MapClaims), false

	switch actor {
	case "user":
		act, ok = utils.JwtParseUser(signed)
	case "admin":
		act, ok = utils.JwtParseAdmin(signed)
	}

	return act, ok
}
