package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
	"time"
)

type DefaultTransactionService struct {
	repository domain.TransactionRepository
}

func (service DefaultTransactionService) NewTransaction(r dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	//Request validation
	validationError := r.Validate()
	if validationError != nil {
		logger.Error("Request validation failed")
		return nil, validationError
	}
	transaction := domain.Transaction{
		TransactionID:   "",
		AccountID:       r.AccountID,
		Amount:          r.Amount,
		TransactionType: r.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	newTransaction, err := service.repository.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	transactionResponseDTO := newTransaction.ToTransactionResponseDTO()
	return &transactionResponseDTO, nil
}

func NewDefaultTransactionService(repository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repository: repository}
}
