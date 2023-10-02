package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
)

// This is the primary port of our hexagonal architecture.
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerWith(customerID string) (*dto.CustomerResponse, *errs.AppError)
}

// Adapter for the primary port
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerWith(customerID string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.Find(customerID)
	if err != nil {
		logger.Error("Error found")
		return nil, err
	}
	customerResponse := customer.ToDTO()
	return &customerResponse, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
