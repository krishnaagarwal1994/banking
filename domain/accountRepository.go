package domain

import "banking/errs"

// AccountRepository PORT (interface)
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FetchDetails(string) (*Account, *errs.AppError)
	CreateTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}
