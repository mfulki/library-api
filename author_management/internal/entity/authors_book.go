package entity

import (
	"author-service/internal/apperror"
	"author-service/pkg/llog"
	"encoding/json"
)

type AuthorsBook struct {
	AuthorId   uint64 `json:"author_id"`
	AuthorName string `json:"author_name"`
}

type AuthorsBooksJson []AuthorsBook

func (j *AuthorsBooksJson) Scan(value interface{}) error {
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
