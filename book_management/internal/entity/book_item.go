package entity

import (
	"book-service/internal/apperror"
	"book-service/pkg/llog"
	"encoding/json"
)

type BookItem struct {
	Id     uint64 `json:"book_item_id"`
	BookId uint64
	Status string `json:"status"`
}
type BookItemJson []BookItem

func (j *BookItemJson) Scan(value interface{}) error {
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
