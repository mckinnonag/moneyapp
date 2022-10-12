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

func (m mockTransactionRepo) CreateTransaction(request api.NewTransactionRequest) error {
	if request.Name == "test Transaction already created" {
		return errors.New("repository - Transaction already exists in database")
	}

	return nil
}

func TestCreateNewTransaction(t *testing.T) {
	mockRepo := mockTransactionRepo{}
	mockTransactionService := api.NewTransactionService(&mockRepo)

	tests := []struct {
		name    string
		request api.NewTransactionRequest
		want    error
	}{
		{
			name: "should create a new Transaction successfully",
			request: api.NewTransactionRequest{
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
			want: nil,
		}, {
			name: "should return an error because of missing id",
			request: api.NewTransactionRequest{
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
			want: errors.New("tx service - transaction id required"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("uid", "123")

			err := mockTransactionService.New(c, test.request)

			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}
