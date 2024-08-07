package usecase

import (
	"book-service/internal/apperror"
	"book-service/internal/database/transaction"
	"book-service/internal/entity"
	"book-service/internal/repository"
	"book-service/pkg/utils"
	"context"
)

type BookUsecase interface {
	GetAllBook(ctx context.Context) (*entity.Books, error)
	GetBook(ctx context.Context, id uint64) (*entity.Book, error)
	UserBorrowBook(ctx context.Context, ids []uint64) error
	UserReturnsBook(ctx context.Context, ids []uint64) error
}

type bookUsecaseImpl struct {
	bookRepository         repository.BookRepository
	bookItemRepository     repository.BookItemRepository
	transactor             transaction.Transactor
	stockJournalRepository repository.StockJournalRepository
}

func NewBookUsecase(bookRepository repository.BookRepository,
	bookItemRepository repository.BookItemRepository,
	transactor transaction.Transactor,
	stockJournalRepository repository.StockJournalRepository) *bookUsecaseImpl {
	return &bookUsecaseImpl{
		bookRepository:         bookRepository,
		bookItemRepository:     bookItemRepository,
		transactor:             transactor,
		stockJournalRepository: stockJournalRepository,
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

func (u *bookUsecaseImpl) UserBorrowBook(ctx context.Context, ids []uint64) error {
	user, ok := utils.CtxGetUser(ctx)
	if !ok {
		return apperror.ErrUnauthorized
	}
	futureStatus := "borrowed"
	currentStatus := "available"

	_, err := u.transactor.WithTransaction(ctx, func(ctx context.Context) (any, error) {
		err := u.bookItemRepository.UpdateStatusBookItems(ctx, futureStatus, currentStatus, ids)
		if err != nil {
			return nil, err
		}

		return nil, u.stockJournalRepository.InsertStockJournal(ctx, futureStatus, ids, user.Id)
	})
	return err
}

func (u *bookUsecaseImpl) UserReturnsBook(ctx context.Context, ids []uint64) error {
	user, ok := utils.CtxGetUser(ctx)
	if !ok {
		return apperror.ErrUnauthorized
	}
	futureStatusBook := "available"
	futureStatus := "returned"
	currentStatus := "borrowed"

	_, err := u.transactor.WithTransaction(ctx, func(ctx context.Context) (any, error) {
		err := u.bookItemRepository.UpdateStatusBookItems(ctx, futureStatusBook, currentStatus, ids)
		if err != nil {
			return nil, err
		}
		return nil, u.stockJournalRepository.InsertStockJournal(ctx, futureStatus, ids, user.Id)
	})
	return err
}
