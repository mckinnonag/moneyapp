package app_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"moneyapp/pkg/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockTransactionService struct{}

func (t *mockTransactionService) New(c *gin.Context, txs []api.Transaction) error {
	_, exists := c.Get("uid")
	if !exists {
		return errors.New("request context does not contain user id claim")
	}

	for _, t := range txs {
		if t.ID == "" {
			return errors.New("tx service - transaction id required")
		}
	}
	return nil
}

func (t *mockTransactionService) Get(c *gin.Context) ([]api.Transaction, error) {
	return []api.Transaction{
			{
				ID:              "",
				ItemID:          "plaidid",
				Category:        nil,
				Location:        "New York",
				Name:            "Shoe Shine LLC",
				Amount:          420.69,
				IsoCurrencyCode: "USD",
				Date:            "2016-06-22 19:10:25-07",
				Pending:         false,
				MerchantName:    "McDonalds",
				PaymentChannel:  "",
				SharedWith:      "",
				SplitAmount:     69.00,
			},
		},
		nil
}

func TestCreateTransactions(t *testing.T) {
	mockServer := newMockServer()

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = &http.Request{
			Header: make(http.Header),
		}
		c.Request.Method = "POST"

		tx := `{
				"transactions": [
					{
						"transaction_id": "456",
						"uid": "278",
						"item_id": "plaidid",
						"category": ["shopping"],
						"location": "New York",
						"name": "Shoe Shine LLC",
						"amount": 420.69,
						"iso_currency_code": "USD",
						"date": "2016-06-22 19:10:25-07",
						"pending": false,
						"merchant_name": "McDonalds",
						"payment_channel": "1",
						"shared_with": "2",
						"split_amount": 69.00
					}
				]
		}`
		jsonBytes := json.RawMessage(tx)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
		c.Set("uid", "123")

		f := mockServer.CreateTransactions()
		f(c)

		var resp gin.H
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

// func TestGetTransactions(t *testing.T) {

// }
