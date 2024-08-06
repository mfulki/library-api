package example

import "context"

type Usecase interface {
	GetOne(context.Context, int) (int, error)
	StoreOne(context.Context, Entity) (int, error)
}

type usecaseImpl struct {
	exampleRepo Repository
}

func NewUsecase(exampleRepo Repository) *usecaseImpl {
	return &usecaseImpl{
		exampleRepo: exampleRepo,
	}
}

func (u *usecaseImpl) GetOne(ctx context.Context, exampleID int) (int, error) {
	return u.exampleRepo.SelectOneByID(ctx, exampleID)
}

func (u *usecaseImpl) StoreOne(ctx context.Context, newExample Entity) (int, error) {
	return u.exampleRepo.InsertOne(ctx, newExample)
}
