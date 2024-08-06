package entity

import (
	"category-service/internal/apperror"
	"category-service/pkg/llog"
	"encoding/json"
)

type BookCategory struct {
	CategoryId   uint64 `json:"category_id"`
	CategoryName string `json:"category_name"`
}
type BookCategoryJson []BookCategory

func (j *BookCategoryJson) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	byteValue, ok := value.([]byte)
	if !ok {
		llog.Error("cannot scan type %T into JSONB", value)
		return apperror.ErrInvalidRequest
	}
	return json.Unmarshal(byteValue, j)
}
