package entity

import "database/sql"

type BookCategory struct {
	Id           uint64
	CategoryId   uint64
	CategoryName string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	DeletedAt    sql.NullTime
}
