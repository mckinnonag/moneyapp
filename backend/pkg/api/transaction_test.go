package api_test

import (
	"errors"
	"moneyapp/pkg/api"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockTransactionRepo struct{}

func (m mockTransactionRepo) CreateTransactions(txs []api.Transaction) error {
	return nil
}

func (m mockTransactionRepo) GetTransactions(uid string) ([]api.Transaction, error) {
	return nil, errors.New("not implemented")
}

func TestCreateNewTransaction(t *testing.T) {
	mockRepo := mockTransactionRepo{}
	mockTransactionService := api.NewTransactionService(&mockRepo)

	tests := []struct {
		name         string
		transactions []api.Transaction
		want         error
	}{
		{
			name: "should create a new Transaction successfully",
			transactions: []api.Transaction{
				{
					ID:              "xyz",
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
			want: nil,
		}, {
			name: "should return an error because of missing id",
			transactions: []api.Transaction{
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
			want: errors.New("tx service - transaction id required"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("uid", "123")

			err := mockTransactionService.New(c, test.transactions)

			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}
