package domain

// This is the secondary port of our hexagonal architecture
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	Find(customerID string) (*Customer, error)
}
