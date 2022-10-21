package api

type NewTransactionsRequest struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID              string   `json:"transaction_id"`
	UID             string   `json:"uid"`
	ItemID          string   `json:"item_id"`
	Category        []string `json:"category"`
	Location        string   `json:"location"`
	Name            string   `json:"name"`
	Amount          float32  `json:"amount"`
	IsoCurrencyCode string   `json:"iso_currency_code"`
	Date            string   `json:"date"`
	Pending         bool     `json:"pending"`
	MerchantName    string   `json:"merchant_name"`
	PaymentChannel  string   `json:"payment_channel"`
	SharedWith      string   `json:"shared_with"`
	SplitAmount     float32  `json:"split_amount"`
}

type NewAccessTokenRequest struct {
	UID         string
	PublicToken string `json:"public_token"`
	AccessToken string
	ItemId      string
}

type GetPlaidTransactionsRequest struct {
	UID       string
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Count     string `json:"count"`
	Offset    string `json:"offset"`
}
