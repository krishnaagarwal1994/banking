package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
)

// This is the primary port of our hexagonal architecture.
type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, error)
	GetCustomerWith(customerID string) (*dto.CustomerResponse, *errs.AppError)
}

// Adapter for the primary port
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		logger.Error("Error found")
		return nil, err
	}
	customerResponses := make([]dto.CustomerResponse, 0)
	for i := 0; i < len(customers); i++ {
		customerResponse := customers[i].ToDTO()
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses, nil
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
