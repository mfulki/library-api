package handler

import (
	pb "category-service/internal/pb/categories"
	"category-service/internal/usecase"
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


