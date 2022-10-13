package app_test

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type mockPlaidService struct{}

func init() {
	gin.SetMode(gin.TestMode)
}

func (m *mockPlaidService) CreateLinkToken(c *gin.Context) (string, error) {
	return "testLinkToken", nil
}

func (m *mockPlaidService) GetAccessToken(c *gin.Context, a NewAccessTokenRequest) (string, string, error) {
	return "testAccessToken", "testItemID", nil
}

func TestCreateLinkToken(t *testing.T) {
	t.Run("Success", func(t *testing.T)) {
		var mockService := mockPlaidService
		
		r := httptest.NewRecorder()
	}
}

func TestGetAccessToken(t *testing.T) {

}