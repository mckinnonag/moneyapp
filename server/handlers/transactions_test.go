package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetSharedTransactions(t *testing.T) {
	mockResponse := `{"transactions":"Welcome to the Tech Company listing API with Golang"}`
	r := SetUpRouter()
	r.GET("/api/private/gettransactions", GetPlaidTransactions)
	req, _ := http.NewRequest("GET", "/api/private/gettransactions", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
