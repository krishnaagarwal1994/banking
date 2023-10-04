package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"strconv"
)

// AccountRepositoryDB Adapter (Implementation confirming to port)
type AccountRepositoryDB struct {
	client *sql.DB
}

func (db AccountRepositoryDB) Save(account Account) (*Account, *errs.AppError) {
	insertQuery := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ? , ?, ?, ?)"
	result, err := db.client.Exec(insertQuery, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error while inserting new account into the DB")
		return nil, errs.NewServerError("Unexpected error from the DB")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Unable to fetch the last inserted ID")
		return nil, errs.NewServerError("Error in fetching the last inserted AccountID")
	}
	account.AccountID = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDB(client *sql.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: client}
}
