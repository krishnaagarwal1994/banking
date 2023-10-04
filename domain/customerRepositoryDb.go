package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Actual DB Adapter for the secondary port i.e CustomerRepository
type CustomerRepositoryDb struct {
	client *sql.DB
}

func (repository CustomerRepositoryDb) FindAll() ([]Customer, error) {
	customers := make([]Customer, 0)

	allCustomers := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, error := repository.client.Query(allCustomers)
	if error != nil {
		logger.Error("Error while querying the customers from the DB")
		return nil, error
	}

	for rows.Next() {
		var c Customer
		error := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if error != nil {
			logger.Error("Error while scanning the customer rows")
			return nil, error
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (repository CustomerRepositoryDb) Find(customerID string) (*Customer, *errs.AppError) {
	if customerID == "" {
		return nil, errs.NewServerError("Customer id is empty")
	}
	findCustomerQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := repository.client.QueryRow(findCustomerQuery, customerID)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug("Customer not found")
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error occured while scanning the customer id " + customerID)
			return nil, errs.NewServerError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(client *sql.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: client}
}
