package models

import (
	"time"
)

type Transaction struct {
	Amount   float32
	Category string
	Date     time.Time
}

var Transactions = []Transaction{
	{Amount: 15.32, Category: "Shopping", Date: time.Now()},
	{Amount: 9.01, Category: "Food", Date: time.Now()},
	{Amount: 126.00, Category: "Bills", Date: time.Now()},
}

func GetAllTransactions(email string) []Transaction {
	return Transactions
}
