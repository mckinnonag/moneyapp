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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockPlaidService struct {
}

func init() {
	gin.SetMode(gin.TestMode)
}

func (m mockPlaidService) CreateLinkToken(c *gin.Context) (string, error) {
	return "testLinkToken", nil
}

func (m mockPlaidService) GetAccessToken(c *gin.Context, a api.NewAccessTokenRequest) (string, string, error) {
	return "testAccessToken", "testItemID", nil
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
		if err != nil {
			panic(err)
		}
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
