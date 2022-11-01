package app

import (
	"encoding/json"
	"io"
	"net/http"

	"moneyapp/pkg/api"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		body, err := io.ReadAll(c.Request.Body)

		if err != nil {
			s.l.Error("handler error: %v", err.Error())
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		req := api.NewTransactionsRequest{}
		err = json.Unmarshal(body, &req)

		if err != nil {
			s.l.Error("handler error: %v", err.Error())
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.transactionService.New(c, req.Transactions)

		if err != nil {
			s.l.Error("service error: %v", err.Error())
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

func (s *Server) GetTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		txs, err := s.transactionService.Get(c)
		if err != nil {
			s.l.Error("service error: %v", err.Error())
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"transactions": txs,
		})
	}
}
