package app

import (
	"log"
	"moneyapp/pkg/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateLinkToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		linkToken, err := s.plaidService.CreateLinkToken(c)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{"link_token": linkToken})
	}
}

func (s *Server) GetAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, exists := c.Get("uid")
		if !exists {
			log.Printf("request context does not contain user id claim")
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		var newAccessToken api.NewAccessTokenRequest
		err := c.ShouldBindJSON(&newAccessToken)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		newAccessToken.UID = uid.(string)

		accessToken, itemID, err := s.plaidService.GetAccessToken(c, newAccessToken)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": accessToken,
			"item_id":      itemID,
		})
	}
}
