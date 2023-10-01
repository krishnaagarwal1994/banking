package domain

import (
	"banking/errs"
	"database/sql"
	"log"
	"time"

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
		log.Print("Error while querying the customers from the DB")
		return nil, error
	}

	for rows.Next() {
		var c Customer
		error := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if error != nil {
			log.Print("Error while scanning the customer rows")
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
			log.Print("Customer not found")
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Print("Error occured while scanning the customer id " + customerID)
			return nil, errs.NewServerError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	db, err := sql.Open("mysql", "root:Gn1d0c@123@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: db}
}
