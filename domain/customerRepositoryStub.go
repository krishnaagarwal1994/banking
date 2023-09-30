package domain

// Adapter for the secondary port
type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Krishna", "Gwalior", "474006", "31-08-1994", "single"},
		{"2", "Madhur", "Morena", "474005", "21-11-1994", "single"},
	}
	return CustomerRepositoryStub{customers: customers}
}
