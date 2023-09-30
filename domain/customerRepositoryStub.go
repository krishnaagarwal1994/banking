package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Krishna", "Gwalior", "474006"},
		{"2", "Madhur", "Morena", "474001"},
	}
	return CustomerRepositoryStub{customers: customers}
}
