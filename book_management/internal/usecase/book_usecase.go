package usecase

import (
	"book-service/internal/repository"
)

type BookUsecase interface {
}

type bookUsecaseImpl struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) *bookUsecaseImpl {
	return &bookUsecaseImpl{
		bookRepository: bookRepository,
	}
}

func (u *bookUsecaseImpl) UserBorrowBook() {

}
