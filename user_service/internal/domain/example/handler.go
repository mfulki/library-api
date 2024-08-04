package example

import (
	"net/http"
	"strconv"

	"user-service/internal/constant"
	"user-service/internal/dto"
	"user-service/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	exampleUsecase Usecase
}

func NewHandler(exampleUsecase Usecase) *Handler {
	return &Handler{
		exampleUsecase: exampleUsecase,
	}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	r := router.Group("/examples")
	r.Get("/:exampleId", h.getOne)
	r.Post("/", h.storeOne)
}

func (h *Handler) getOne(ctx *fiber.Ctx) error {
	exampleID, err := strconv.Atoi(ctx.Params("exampleId"))
	if err != nil {
		return err
	}

	result, err := h.exampleUsecase.GetOne(ctx.Context(), exampleID)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Message: constant.DataRetrievedMsg,
		Data:    result,
	})
}

func (h *Handler) storeOne(ctx *fiber.Ctx) error {
	reqBody := new(StoreRequest)
	if err := validate.BodyJSON(ctx, reqBody); err != nil {
		return err
	}

	result, err := h.exampleUsecase.StoreOne(ctx.Context(), reqBody.Example())
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(dto.Response{
		Message: constant.DataCreatedMsg,
		Data:    result,
	})
}
