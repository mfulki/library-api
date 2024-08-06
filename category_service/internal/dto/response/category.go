package response

import (
	"category-service/internal/entity"
	pb "category-service/internal/pb/categories"
)

func NewBookCategoryResp(mapBC map[uint64]entity.BookCategoryJson) *pb.BookCategoriesMap {
	mapBookCategory := make(map[uint64]*pb.BookCategories, 0)
	for key, val := range mapBC {
		sliceBookCategory := []*pb.BookCategory{}
		for _, bookCategory := range val {
			pbBookCategory := pb.BookCategory{ CategoryId: bookCategory.CategoryId, CategoryName: bookCategory.CategoryName}
			sliceBookCategory = append(sliceBookCategory, &pbBookCategory)
		}
		BookCategories := pb.BookCategories{BookCategoriesList: sliceBookCategory}
		mapBookCategory[key] = &BookCategories
	}
	return &pb.BookCategoriesMap{
		BookCategoriesMap: mapBookCategory,
	}
}
