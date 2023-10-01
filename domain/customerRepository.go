package domain

import "banking/errs"

// This is the secondary port of our hexagonal architecture
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	Find(customerID string) (*Customer, *errs.AppError)
}
