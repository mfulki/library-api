package usecase

import (
	"context"
	"user-service/internal/constant"
	"user-service/internal/entity"
	"user-service/internal/repository"
	"user-service/pkg/utils"
)

type AuthUsecase interface {
	Register(ctx context.Context, user entity.User) error
}

type authUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) *authUsecaseImpl {
	return &authUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *authUsecaseImpl) Register(ctx context.Context, user entity.User) error {
	_, err := u.userRepository.SelectOneByEmail(ctx, user)
	if err != nil {
		return err
	}
	user.PhotoURL = constant.DefaultPhotoURL
	user.Password, err = utils.HashPasswordDefault(user.Password)
	if err != nil {
		return err
	}
	err = u.userRepository.RegisterUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
