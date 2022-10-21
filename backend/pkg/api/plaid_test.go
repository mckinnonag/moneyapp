package api_test

import (
	"fmt"
	"moneyapp/pkg/api"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var ACCESS_TOKEN string

type mockPlaidRepo struct{}

func (m *mockPlaidRepo) CreateAccessToken(request *api.NewAccessTokenRequest) error {
	return nil
}

func (m *mockPlaidRepo) GetAccessTokens(string) ([]string, error) {
	var ret []string
	ret = append(ret, ACCESS_TOKEN)
	return ret, nil
}

var mockPlaidConfig api.PlaidConfig

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	gin.SetMode(gin.TestMode)

	mockPlaidConfig.APP_NAME = "test runner"
	mockPlaidConfig.PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID")
	mockPlaidConfig.PLAID_SECRET = os.Getenv("PLAID_SECRET")
	mockPlaidConfig.PLAID_REDIRECT_URI = ""
	mockPlaidConfig.PLAID_ENV = "sandbox"
	mockPlaidConfig.PLAID_PRODUCTS = "transactions"
	ACCESS_TOKEN = os.Getenv("PLAID_SANDBOX_ACCESS_TOKEN")
}

func TestCreateLinkToken(t *testing.T) {
	mockRepo := mockPlaidRepo{}
	mockPlaidService := api.NewPlaidService(&mockPlaidConfig, &mockRepo)

	t.Run("should create a new linkToken", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("uid", "123")

		linkToken, err := mockPlaidService.CreateLinkToken(c)

		assert.Nil(t, err)

		expected := 49
		assert.Lenf(t, linkToken, expected, "expected token with length %d, got %d", expected, len(linkToken))
	})
}

func TestGetAccessToken(t *testing.T) {
	mockRepo := mockPlaidRepo{}
	mockPlaidService := api.NewPlaidService(&mockPlaidConfig, &mockRepo)

	publicToken, err := api.CreatePublicToken()
	assert.Nil(t, err)

	t.Run("should create a new access token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("uid", "123")

		request := api.NewAccessTokenRequest{
			PublicToken: publicToken,
		}
		accessToken, itemID, err := mockPlaidService.GetAccessToken(c, &request)

		assert.Nil(t, err)

		expected := 51
		assert.Lenf(t, accessToken, expected, "expected token with length %d, got %d", expected, len(accessToken))

		expected = 37
		assert.Lenf(t, itemID, expected, "expected token with length %d, got %d", expected, len(itemID))
	})

	t.Run("require public token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("uid", "123")

		request := api.NewAccessTokenRequest{}
		_, _, err := mockPlaidService.GetAccessToken(c, &request)

		assert.NotNil(t, err)
	})
}

func TestGetPlaidTransactions(t *testing.T) {
	mockRepo := mockPlaidRepo{}
	mockPlaidService := api.NewPlaidService(&mockPlaidConfig, &mockRepo)

	t.Run("should fetch user's transactions", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("uid", "123")

		const iso8601TimeFormat = "2006-01-02"
		startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
		endDate := time.Now().Format(iso8601TimeFormat)

		request := api.GetPlaidTransactionsRequest{
			UID:       "123",
			StartDate: startDate,
			EndDate:   endDate,
			Count:     "100",
			Offset:    "0",
		}

		txs, err := mockPlaidService.GetPlaidTransactions(c, request)
		assert.Nil(t, err)
		assert.Len(t, txs, 100)
	})
}
