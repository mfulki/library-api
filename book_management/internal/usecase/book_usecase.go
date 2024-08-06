package usecase

import (
	"book-service/internal/entity"
	"book-service/internal/repository"
	"context"
)

type BookUsecase interface {
	GetAllBook(ctx context.Context) (*entity.Books, error)
	GetBook(ctx context.Context, id uint64) (*entity.Book, error)
}

type bookUsecaseImpl struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) *bookUsecaseImpl {
	return &bookUsecaseImpl{
		bookRepository: bookRepository,
	}
}

func (u *bookUsecaseImpl) GetAllBook(ctx context.Context) (*entity.Books, error) {
	books, err := u.bookRepository.GetAllBook(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (u *bookUsecaseImpl) GetBook(ctx context.Context, id uint64) (*entity.Book, error) {
	book, err := u.bookRepository.GetOneBook(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
func (u *bookUsecaseImpl) UserBorrowBook() {
	
}
