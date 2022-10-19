package api

type NewTransactionRequest struct {
	ID              string   `json:"id"`
	UID             string   `json:"uid"`
	ItemID          string   `json:"itemid"`
	Category        []string `json:"category"`
	Location        string   `json:"location"`
	Name            string   `json:"name"`
	Amount          float32  `json:"amount"`
	IsoCurrencyCode string   `json:"isocurrencycode"`
	Date            string   `json:"date"`
	Pending         bool     `json:"pending"`
	MerchantName    string   `json:"merchantname"`
	PaymentChannel  string   `json:"paymentchannel"`
	SharedWith      string   `json:"sharedwith"`
	SplitAmount     float32  `json:"splitamount"`
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
