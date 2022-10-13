package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetContact() gin.HandlerFunc {
	return func(c *gin.Context) {
		val, err := s.contactService.GetContact(c)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{"key": val})
	}
}

func (s *Server) SetContact() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (s *Server) DeleteContact() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
