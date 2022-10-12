package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateLinkToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		linkToken, err := s.plaidService.CreateLinkToken(c)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{"link_token": linkToken})
	}
}
