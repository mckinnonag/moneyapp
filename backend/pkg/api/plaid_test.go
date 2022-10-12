package api_test

import (
	"fmt"
	"moneyapp/pkg/api"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type mockPlaidRepo struct{}

var mockPlaidConfig api.PlaidConfig

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	mockPlaidConfig.APP_NAME = "test runner"
	mockPlaidConfig.PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID")
	mockPlaidConfig.PLAID_SECRET = os.Getenv("PLAID_SECRET")
	mockPlaidConfig.PLAID_REDIRECT_URI = ""
	mockPlaidConfig.PLAID_ENV = "sandbox"
}

func TestCreateLinkToken(t *testing.T) {
	mockRepo := mockPlaidRepo{}
	mockPlaidService := api.NewPlaidService(&mockPlaidConfig, &mockRepo)

	tests := []struct {
		name string
		want string
		err  error
	}{
		{
			name: "should create a new linkToken",
			want: "string",
			err:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("uid", "123")

			linkToken, err := mockPlaidService.CreateLinkToken(c)

			assert.Nil(t, err)
			assert.Lenf(t, linkToken, 49, "expected token with length %d, got %d", 49, len(linkToken))
		})
	}
}
