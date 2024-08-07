package entity

import (
	"database/sql"
)

type Category struct {
	Id        uint64
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
