package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"strconv"
)

type TransactionRepositoryDB struct {
	client *sql.DB
}

func (db TransactionRepositoryDB) CreateTransaction(t Transaction) (*Transaction, *errs.AppError) {
	sqlQuery := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	result, err := db.client.Exec(sqlQuery, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error creating the transaction record")
		return nil, errs.NewUnexpectedError("Error creating the transaction record")
	}
	transactionID, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error occured while fetching the last inserted transaction ID")
		return nil, errs.NewUnexpectedError("Error occured while fetching the last inserted transaction ID")
	}
	t.TransactionID = strconv.FormatInt(transactionID, 10)
	return &t, nil
}

func NewTransactionRepository(client *sql.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{client: client}
}
