package request

import (
	"time"
	"user-service/internal/constant"
	"user-service/internal/entity"
)

type UserRegister struct {
	Name        string `json:"name" validate:"required,min=2"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	DateOfBirth string `json:"date_of_birth" validate:"required,date"`
	Gender      string `json:"gender" validate:"required,oneof=male female"`
}

func (req *UserRegister) Auth() entity.User {
	dateOfBirth, _ := time.Parse(constant.DateFormat, req.DateOfBirth)

	return entity.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    &req.Password,
		DateOfBirth: dateOfBirth,
		Gender:      req.Gender,
	}
}
