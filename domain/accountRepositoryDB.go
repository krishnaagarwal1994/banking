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

func (db AccountRepositoryDB) FetchDetails(accountID string) (*Account, *errs.AppError) {
	fetchSqlQuery := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"
	row := db.client.QueryRow(fetchSqlQuery, accountID)
	var account Account
	err := row.Scan(&account.AccountID, &account.CustomerID, &account.OpeningDate, &account.AccountType, &account.Amount, &account.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug("Bank Account not found")
			return nil, errs.NewNotFoundError("Bank Account not found")
		} else {
			logger.Error("Error fetching account details from DB")
			return nil, errs.NewUnexpectedError("Error fetching account details from DB")
		}
	} else {
		return &account, nil
	}
}

func (db AccountRepositoryDB) CreateTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// Starting the database transaction block
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for the bank account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	// Inserting Bank account transaction
	var result sql.Result
	insertSqlQuery := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	result, _ = tx.Exec(insertSqlQuery, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)

	// Updating account balance
	var updateSqlQuery string
	if t.IsWithdrawal() {
		updateSqlQuery = "UPDATE accounts SET amount =  amount - ? where account_id = ?"
	} else {
		updateSqlQuery = "UPDATE accounts SET amount =  amount + ? where account_id = ?"
	}
	_, err = tx.Exec(updateSqlQuery, t.Amount, t.AccountID)

	//In case of error, rollback the transaction, and changes from both the tables will be reverted.
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction")
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	// Commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error when committing the transaction")
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	// Getting the last inserted transaction id
	transactionID, transactionError := result.LastInsertId()
	if transactionError != nil {
		logger.Error("Error while fetching the last transaction ID")
		return nil, errs.NewUnexpectedError("Database error")
	}

	// Getting the latest amount information from the account table
	account, error := db.FetchDetails(t.AccountID)
	if error != nil {
		logger.Error("Failed to fetch the account details post transaction")
		return nil, errs.NewUnexpectedError("Database error")
	}
	t.Amount = account.Amount
	t.TransactionID = strconv.FormatInt(transactionID, 10)
	return &t, nil
}

func NewAccountRepositoryDB(client *sql.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: client}
}
