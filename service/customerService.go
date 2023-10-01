package service

import "banking/domain"

// This is the primary port of our hexagonal architecture.
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerWith(customerID string) (*domain.Customer, error)
}

// Adapter for the primary port
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerWith(customerID string) (*domain.Customer, error) {
	return s.repo.Find(customerID)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
