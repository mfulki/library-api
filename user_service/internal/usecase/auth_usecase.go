package usecase

import (
	"context"
	"user-service/internal/apperror"
	"user-service/internal/constant"
	"user-service/internal/entity"
	"user-service/internal/repository"
	"user-service/pkg/utils"
)

type AuthUsecase interface {
	Register(ctx context.Context, user entity.User) error
	Login(ctx context.Context, user entity.User) (string, error)
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

func (u *authUsecaseImpl) Login(ctx context.Context, user entity.User) (string, error) {
	result, err := u.userRepository.SelectOneByEmail(ctx, user)
	if err != nil {
		return "", err
	}
	if !utils.HashCompareDefault(*user.Password, result.Password) {
		return "", apperror.ErrInternalServer
	}

	jwtData := map[string]any{
		"ID":    user.ID,
		"Email": user.Email,
	}

	token, err := utils.JwtGenerateUser(jwtData)
	if err != nil {
		return "", err
	}
	return token, nil

}
