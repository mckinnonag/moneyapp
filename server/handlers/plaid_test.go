package handlers

import (
	"fmt"
	"net/http/httptest"
	"testing"

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

// func TestLinkTokenCreate(t *testing.T) {
// 	token, err := linkTokenCreate(nil)
// 	assert.NoError(t, err)
// 	fmt.Println(token)
// }

func TestCreateLinkToken(t *testing.T) {
	// router := gin.Default()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	CreateLinkToken(c)
	assert.Equal(t, 200, w.Code)
}
