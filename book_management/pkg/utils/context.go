package utils

import (
	"book-service/internal/apperror"
	"book-service/internal/constant"
	"book-service/internal/entity"
	"context"

	"github.com/sirupsen/logrus"
)

func CtxGetUser(ctx context.Context) (*entity.User, bool) {
	userMap, ok := getDetailActor(ctx, constant.User)
	if !ok {
		return nil, false
	}

	return &entity.User{
		Id:    uint64(userMap["ID"].(float64)),
		Email: userMap["Email"].(string),
	}, true
}

// func CtxGetAdmin(ctx context.Context) (*entity.Admin, bool) {
// 	adminMap, ok := getDetailActor(ctx, constant.Admin)
// 	if !ok {
// 		return nil, false
// 	}

//		return &entity.Admin{
//			ID:    uint(adminMap["ID"].(float64)),
//			Email: adminMap["Email"].(string),
//		}, true
//	}
func getDetailActor(ctx context.Context, actor string) (map[string]any, bool) {
	val := ctx.Value(constant.UserContext)

	act, ok := val.(map[string]any)
	if !ok {
		logrus.Error(apperror.ErrAssertingAny)
		return nil, false
	}

	if act["role"] != actor {
		return nil, false
	}

	return act, true
}
