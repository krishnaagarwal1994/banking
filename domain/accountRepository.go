package domain

import "banking/errs"

// AccountRepository PORT (interface)
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
