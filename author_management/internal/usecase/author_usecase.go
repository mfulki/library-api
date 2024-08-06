package usecase

import (
	"author-service/internal/entity"
	"author-service/internal/repository"
	"context"
)

type AuthorUsecase interface {
	GetOneAuthor(ctx context.Context, author entity.Author) (*entity.Author, error)
	GetAllAuthor(ctx context.Context) ([]entity.Author, error)
	InsertAuthor(ctx context.Context, authors []entity.Author) error
	GetSomeAuthor(ctx context.Context, ids []uint64) ([]entity.Author, error)
	DeleteOneAuthor(ctx context.Context, id uint64) error
}

type authorUsecaseImpl struct {
	authorRepository repository.AuthorRepository
}

func NewAuthorUsecase(authorRepository repository.AuthorRepository) *authorUsecaseImpl {
	return &authorUsecaseImpl{
		authorRepository: authorRepository,
	}
}

func (u *authorUsecaseImpl) GetAllAuthor(ctx context.Context) ([]entity.Author, error) {
	return u.authorRepository.GetAllAuthor(ctx)
}

func (u *authorUsecaseImpl) GetOneAuthor(ctx context.Context, author entity.Author) (*entity.Author, error) {
	return u.authorRepository.GetOneAuthor(ctx, author)
}

func (u *authorUsecaseImpl) InsertAuthor(ctx context.Context, authors []entity.Author) error {
	return u.authorRepository.InsertAuthor(ctx, authors)
}

func (u *authorUsecaseImpl) GetSomeAuthor(ctx context.Context, ids []uint64) ([]entity.Author, error) {
	return u.authorRepository.GetSomeAuthor(ctx, ids)
}

func (u *authorUsecaseImpl) DeleteOneAuthor(ctx context.Context, id uint64) error {
	return u.authorRepository.DeleteOneAuthor(ctx, id)
}
