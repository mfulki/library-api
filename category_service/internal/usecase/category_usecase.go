package usecase

import (
	"category-service/internal/repository"
)

type CategoryUsecase interface {
}

type categoryUsecaseImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepository repository.CategoryRepository) *categoryUsecaseImpl {
	return &categoryUsecaseImpl{
		categoryRepository: categoryRepository,
	}
}
