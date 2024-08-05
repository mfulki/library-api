package entity

import "time"

type User struct {
	ID          uint
	Name        string
	Email       string
	Password    *string
	DateOfBirth time.Time
	Gender      string
	PhotoURL    string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
