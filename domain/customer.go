package domain

type Customer struct {
	Id      string `json:"id" xml:"id"`
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

// This is the secondary port of our hexagonal architecture
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
