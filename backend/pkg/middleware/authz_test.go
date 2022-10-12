//go:build integration
// +build integration

package middleware

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"server/auth"
	"server/handlers"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}
}

func TestAuthzNoHeader(t *testing.T) {
	router := gin.Default()
	router.Use(Authz())

	uri := "/api/private/transactions"

	router.GET(uri, handlers.GetPlaidTransactions)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 403, w.Code)
}

func TestAuthzInvalidTokenFormat(t *testing.T) {
	router := gin.Default()
	router.Use(Authz())

	uri := "/api/private/transactions"

	router.GET(uri, handlers.GetPlaidTransactions)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "test")

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestAuthzInvalidToken(t *testing.T) {
	invalidToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	router := gin.Default()
	router.Use(Authz())

	uri := "/api/private/transactions"

	router.GET(uri, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", invalidToken)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

// Generates a random password for testing purposes
func randomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func TestValidToken(t *testing.T) {
	router := gin.Default()
	router.Use(Authz())

	uri := "/api/private/test"

	router.GET(uri, handlers.Test)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Generate random password
	password := randomPassword(16)

	// Create test firebase user
	user, err := auth.CreateUser(c, "user@example.com", password, "Test User")
	assert.NoError(t, err)

	// Generate token
	token, err := auth.CreateCustomToken(c, user.UID)
	assert.NoError(t, err)

	// Exchange custom token for ID token
	idToken, err := auth.ExchangeCustomTokenForIDToken(token, os.Getenv("FIREBASE_API_KEY"))
	assert.NoError(t, err)

	// Send request
	req, err := http.NewRequest("GET", uri, nil)
	assert.NoError(t, err)
	req.Header.Add("Authorization", idToken)

	router.ServeHTTP(w, req)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

	// Remove user
	err = auth.DeleteUser(c, user.UID)
	assert.NoError(t, err)
}
