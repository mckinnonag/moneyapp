package models

type Transaction struct {
	ID              string
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

func GetAllTransactions(email string) (result []Transaction, err error) {
	tx := Transaction{
		ID:           "1",
		MerchantName: "McDonalds",
		Amount:       69.420,
	}
	result = append(result, tx)
	return result, nil
}

func UpdateUserTransactions(email string, transactions []Transaction) error {
	for _, tx := range transactions {
		sqlStatement := `
			INSERT INTO transactions (tx_id, plaid_id, user_id, amount, merchant_name)
			VALUES ($1, $2)`
		_, err := DB.Exec(sqlStatement, tx.ID, tx.)
		if err != nil {
			return err
		}
	}
}