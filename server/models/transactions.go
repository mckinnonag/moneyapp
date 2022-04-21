package models

import (
	"database/sql"
	"errors"
)

type SharedTx struct {
	SharedTxID int64   // Primary key
	TxID       int64   // Foreign Key (transactions)
	UserID     string  // Owner
	SharedWith string  // Recipient
	Amount     float32 // Amount owed
}

type Transaction struct {
	ID              int64
	ItemID          string
	Category        []string
	Location        string
	Name            string
	Amount          float32
	IsoCurrencyCode string
	Date            string
	Pending         bool
	MerchantName    string
	PaymentChannel  string
}

func getCurrencyCode(code string) (currency_id string, err error) {
	// Helper function. Accepts 3 char currency code; returns ID in database
	sqlStatement := `SELECT currency_id, code FROM users WHERE code=$1;`
	row := DB.QueryRow(sqlStatement, code)
	switch err = row.Scan(&currency_id, &code); err {
	case sql.ErrNoRows:
		return "", errors.New("code does not exist")
	case nil:
		return currency_id, nil
	default:
		return "", err
	}
}

func GetSharedTransactions(email string) (result []Transaction, err error) {
	tx := Transaction{
		ID:           "1",
		MerchantName: "McDonalds",
		Amount:       69.420,
	}
	result = append(result, tx)
	return result, nil
}

func ShareTransactions(email string, transactions []Transaction) error {
	user_id, err := lookupUser(email)
	if err != nil {
		return err
	}

	for _, tx := range transactions {
		iso_currency_code, err := getCurrencyCode(tx.IsoCurrencyCode)
		if err != nil {
			return err
		}
		sqlStatement := `
			INSERT INTO transactions (tx_id, plaid_id, plaid_item_id, user_id, category, location, tx_name, amount, iso_currency_code, tx_date, pending, merchant_name, payment_channel)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
		_, err = DB.Exec(sqlStatement, tx.ID, tx.ID, tx.ItemID, user_id, tx.Category, tx.Location, tx.Name, tx.Amount, iso_currency_code, tx.Date, tx.Pending, tx.MerchantName, tx.PaymentChannel)
		if err != nil {
			return err
		}
	}
	return nil
}
