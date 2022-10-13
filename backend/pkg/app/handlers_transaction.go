package app

import (
	"log"
	"net/http"

	"moneyapp/pkg/api"

	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newTransaction api.NewTransactionRequest

		err := c.ShouldBindJSON(&newTransaction)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.transactionService.New(c, newTransaction)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new user created",
		}

		c.JSON(http.StatusOK, response)
	}
}