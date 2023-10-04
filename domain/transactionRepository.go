package domain

import "banking/errs"

type TransactionRepository interface {
	CreateTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}
