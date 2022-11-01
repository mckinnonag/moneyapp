package app

import (
	"moneyapp/pkg/api"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateLinkToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		linkToken, err := s.plaidService.CreateLinkToken(c)
		if err != nil {
			s.l.Error("handler error: %v", err.Error())
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		if linkToken == "" {
			s.l.Error("returned linkToken is null")
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{"link_token": linkToken})
	}
}

func (s *Server) GetAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, exists := c.Get("uid")
		if !exists {
			s.l.Warn("request context does not contain user id claim")
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		var newAccessToken api.NewAccessTokenRequest
		err := c.ShouldBindJSON(&newAccessToken)
		if err != nil {
			s.l.Error("handler error: %v", err.Error())
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		newAccessToken.UID = uid.(string)

		accessToken, itemID, err := s.plaidService.GetAccessToken(c, &newAccessToken)
		if err != nil {
			s.l.Error("service error: %v", err.Error())
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": accessToken,
			"item_id":      itemID,
		})
	}
}

func (s *Server) GetPlaidTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, exists := c.Get("uid")
		if !exists {
			s.l.Warn("request context does not contain user id claim")
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		const iso8601TimeFormat = "2006-01-02"
		startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
		endDate := time.Now().Format(iso8601TimeFormat)

		// newTransactionRequest := api.GetPlaidTransactionsRequest{
		// 	StartDate: c.DefaultQuery("startdate", startDate),
		// 	EndDate:   c.DefaultQuery("enddate", endDate),
		// 	Count:     c.DefaultQuery("count", "100"),
		// 	Offset:    c.DefaultQuery("offset", "0"),
		// 	UID:       uid.(string),
		// }

		newTransactionRequest := api.GetPlaidTransactionsRequest{
			StartDate: startDate,
			EndDate:   endDate,
			Count:     "100",
			Offset:    "0",
			UID:       uid.(string),
		}

		txs, err := s.plaidService.GetPlaidTransactions(c, newTransactionRequest)
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
