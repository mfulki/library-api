package usecase

import (
	"book-service/internal/apperror"
	"book-service/internal/constant"
	"book-service/internal/database/transaction"
	"book-service/internal/entity"
	"book-service/internal/repository"
	"book-service/pkg/utils"
	"context"
	"fmt"
)

type BookUsecase interface {
	GetAllBook(ctx context.Context) (*entity.Books, error)
	GetBook(ctx context.Context, id uint64) (*entity.Book, error)
	UserBorrowBook(ctx context.Context, ids []uint64) (*string, error)
	UserReturnsBook(ctx context.Context, ids []uint64) (*string, error)
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

func (u *bookUsecaseImpl) UserBorrowBook(ctx context.Context, ids []uint64) (*string,error) {
	user, ok := utils.CtxGetUser(ctx)
	if !ok {
		return nil, apperror.ErrUnauthorized
	}
	futureStatus := "borrowed"
	currentStatus := "available"

	resultTx, err := u.transactor.WithTransaction(ctx, func(ctx context.Context) (any, error) {
		mapStatus, err := u.bookItemRepository.LockRow(ctx, ids)
		if err != nil {
			return nil, err
		}
		if mapStatus[currentStatus] == nil {
			return nil, apperror.ErrInvalidRequest
		}
		err = u.bookItemRepository.UpdateStatusBookItems(ctx, futureStatus, currentStatus, mapStatus[currentStatus])
		if err != nil {
			return nil, err
		}
		fmt.Println(err)

		return mapStatus, u.stockJournalRepository.InsertStockJournal(ctx, futureStatus, ids, user.Id)
	})
	if err!=nil{
		return nil,err
	}
	result := resultTx.(map[string][]uint64)
	fmt.Println(result)
	var message string
	if result[currentStatus] != nil {
		message = fmt.Sprintf("Id:%v, %s.", result[currentStatus], constant.BookSuccessReturnMsg)
	}
	for key, value := range result {
		if key != currentStatus {
			message = message + fmt.Sprintf("id %v, cannot borrow book. because status is %s.", value, key)
		}
	}
	return &message,nil
}

func (u *bookUsecaseImpl) UserReturnsBook(ctx context.Context, ids []uint64) (*string, error) {
	user, ok := utils.CtxGetUser(ctx)
	if !ok {
		return nil, apperror.ErrUnauthorized
	}
	futureStatusBook := "available"
	futureStatus := "returned"
	currentStatus := "borrowed"

	resultTx, err := u.transactor.WithTransaction(ctx, func(ctx context.Context) (any, error) {
		mapStatus, err := u.bookItemRepository.LockRow(ctx, ids)
		if err != nil {
			return nil, err
		}
		if mapStatus[currentStatus] == nil {
			return nil, apperror.ErrInvalidRequest
		}
		if err=u.stockJournalRepository.ReturnUpdate(ctx,user.Id,ids);err!=nil{
			return nil,err
		}
		if err = u.bookItemRepository.UpdateStatusBookItems(ctx, futureStatusBook, currentStatus, ids);err!=nil{
			return nil,err
		}
		return mapStatus, u.stockJournalRepository.InsertStockJournal(ctx, futureStatus, ids, user.Id)
	})
	if err != nil {
		return nil, err
	}
	result := resultTx.(map[string][]uint64)
	fmt.Println(result)
	var message string
	if result[currentStatus] != nil {
		message = fmt.Sprintf("Id:%v, %s.", result[currentStatus], constant.BookSuccessReturnMsg)
	}
	for key, value := range result {
		if key != currentStatus {
			message = message + fmt.Sprintf("id %v, cannot return book that you not borrow. because status is %s.", value, key)
		}
	}
	return &message, nil
}
