package handler

import (
	"category-service/internal/dto/response"
	pb "category-service/internal/pb/categories"
	"category-service/internal/usecase"
	"context"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		categoryUsecase: categoryUsecase,
	}
}

func (h *CategoryHandler) GetSomeBookCategories(ctx context.Context, in *pb.Ids) (*pb.BookCategoriesMap, error) {
	res, err := h.categoryUsecase.GetSomeBookCategoriesGetSomeBookCategories(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return response.NewBookCategoryResp(res), nil
}
