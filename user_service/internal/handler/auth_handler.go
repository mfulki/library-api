package handler

import (
	"user-service/internal/constant"
	"user-service/internal/dto"
	"user-service/internal/dto/request"
	"user-service/internal/usecase"
	"user-service/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	body := new(request.UserRegister)

	if err := validate.BodyJSON(ctx, body); err != nil {
		return err
	}

	err := h.authUsecase.Register(ctx.Context(), body.Auth())
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.Response{
		Message: constant.RegisterSuccessMsg,
	})

}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	body := new(request.UserLogin)

	if err := validate.BodyJSON(ctx, body); err != nil {
		return err
	}

	token, err := h.authUsecase.Login(ctx.Context(), body.Auth())
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Message: constant.LoginPassedMsg,
		Data:    token,
	})

}
