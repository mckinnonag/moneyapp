package app_test

import (
	"bytes"
	"encoding/json"
	"io"
	"moneyapp/pkg/api"
	"moneyapp/pkg/app"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/plaid"
	"github.com/stretchr/testify/assert"
)

type mockPlaidService struct{}

func init() {
	gin.SetMode(gin.TestMode)
}

func (m mockPlaidService) CreateLinkToken(c *gin.Context) (string, error) {
	return "testLinkToken", nil
}

func (m mockPlaidService) GetAccessToken(c *gin.Context, a api.NewAccessTokenRequest) (string, string, error) {
	return "testAccessToken", "testItemID", nil
}

func (m mockPlaidService) GetTransactions(r api.GetPlaidTransactionsRequest) ([]plaid.Transaction, error) {
	var ret []plaid.Transaction
	ret = append(ret, plaid.Transaction{
		Name:   "Test Transaction",
		Amount: 69.42,
	})
	return ret, nil
}

func newMockServer() *app.Server {
	var mockPlaidService mockPlaidService
	r := gin.New()

	return app.NewServer(
		r,
		nil,
		mockPlaidService,
	)
}

func TestCreateLinkToken(t *testing.T) {
	mockServer := newMockServer()

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		f := mockServer.CreateLinkToken()
		f(c)

		var resp gin.H
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Nil(t, err)

		expected := "testLinkToken"
		assert.Equal(t, expected, resp["link_token"])
	})
}

func TestGetAccessToken(t *testing.T) {
	mockServer := newMockServer()

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = &http.Request{
			Header: make(http.Header),
		}

		c.Set("uid", "123")
		c.Request.Method = "POST"
		jsonBytes, err := json.Marshal(map[string]string{"public_token": "testPublicToken"})
		assert.Nil(t, err)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))

		f := mockServer.GetAccessToken()
		f(c)

		var resp gin.H
		err = json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "testAccessToken", resp["access_token"])
		assert.Equal(t, "testItemID", resp["item_id"])
	})
	t.Run("401 no uid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		f := mockServer.GetAccessToken()
		f(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
	t.Run("400 no public token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Set("uid", "123")

		f := mockServer.GetAccessToken()
		f(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetPlaidTransactions(t *testing.T) {
	mockServer := newMockServer()

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = &http.Request{
			Header: make(http.Header),
		}

		c.Set("uid", "123")

		const iso8601TimeFormat = "2006-01-02"
		startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
		endDate := time.Now().Format(iso8601TimeFormat)

		c.Request.Method = "POST"
		jsonBytes, err := json.Marshal(map[string]string{
			"start_date": startDate,
			"end_date":   endDate,
			"count":      "1",
			"offset":     "0",
		})
		assert.Nil(t, err)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))

		f := mockServer.GetPlaidTransactions()
		f(c)

		var resp gin.H
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		r := resp["transactions"]
		assert.Len(t, r, 1)

		// Get the first transaction in the array
		first := r.([]interface{})[0].(map[string]interface{})

		assert.Equal(t, "Test Transaction", first["name"])
		assert.Equal(t, 69.42, first["amount"])
	})
}
