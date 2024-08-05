package entity

import (
	"database/sql"
)

type Book struct {
	Id          uint64
	Title       string
	ISBN        string
	Description string
	AuthorId    []int64
	CategoryId  []int64
	BookItemId  []int64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	DeletedAt   sql.NullTime
	BookItem    BookItemJson
}
