package usecase

import (
	"category-service/internal/entity"
	"category-service/internal/repository"
	"context"
)

type CategoryUsecase interface {
	GetSomeBookCategoriesGetSomeBookCategories(ctx context.Context, ids []uint64) (map[uint64]entity.BookCategoryJson, error)
}

type categoryUsecaseImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepository repository.CategoryRepository) *categoryUsecaseImpl {
	return &categoryUsecaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (u *categoryUsecaseImpl) GetSomeBookCategoriesGetSomeBookCategories(ctx context.Context, ids []uint64) (map[uint64]entity.BookCategoryJson, error) {
	return u.categoryRepository.GetSomeBookCategories(ctx, ids)
}
