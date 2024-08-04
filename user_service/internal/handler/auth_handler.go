package handler

import (
	"net/http"
	"user-service/internal/constant"
	"user-service/internal/dto"
	"user-service/internal/dto/request"
	"user-service/internal/usecase"
	"user-service/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
}

type authHandlerImpl struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *authHandlerImpl {
	return &authHandlerImpl{
		authUsecase: authUsecase,
	}
}

func (h *authHandlerImpl) Register(ctx *fiber.Ctx) error {
	body := new(request.UserRegister)

	if err := validate.BodyJSON(ctx, body); err != nil {
		return err
	}

	err := h.authUsecase.Register(ctx.Context(), body.Auth())
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusCreated).JSON(dto.Response{
		Message: constant.RegisterSuccessMsg,
	})

}
