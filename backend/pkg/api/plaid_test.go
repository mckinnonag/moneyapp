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

func (m mockPlaidRepo) CreateAccessToken(request api.NewAccessTokenRequest) error {
	return nil
}

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
	mockPlaidConfig.PLAID_PRODUCTS = "transactions"
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

			expected := 49
			assert.Lenf(t, linkToken, expected, "expected token with length %d, got %d", expected, len(linkToken))
		})
	}
}

func TestGetAccessToken(t *testing.T) {
	mockRepo := mockPlaidRepo{}
	mockPlaidService := api.NewPlaidService(&mockPlaidConfig, &mockRepo)

	publicToken, err := api.CreatePublicToken()
	assert.Nil(t, err)

	tests := []struct {
		name    string
		request api.NewAccessTokenRequest
		want    api.NewAccessTokenResponse
		err     error
	}{
		{
			name: "should create a new access token",
			request: api.NewAccessTokenRequest{
				PublicToken: publicToken,
			},
			want: api.NewAccessTokenResponse{},
			err:  nil,
		},
		// {
		// 	name:    "should be rejecting for missing a public token",
		// 	request: api.NewAccessTokenRequest{},
		// 	want:    api.NewAccessTokenResponse{},
		// 	err:     errors.New("missing public token"),
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("uid", "123")

			accessToken, itemID, err := mockPlaidService.GetAccessToken(c, test.request)

			assert.Nil(t, err)

			expected := 51
			assert.Lenf(t, accessToken, expected, "expected token with length %d, got %d", expected, len(accessToken))

			expected = 37
			assert.Lenf(t, itemID, expected, "expected token with length %d, got %d", expected, len(itemID))
		})
	}
}
