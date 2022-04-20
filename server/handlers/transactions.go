package handlers

import (
	models "server/models"

	"github.com/gin-gonic/gin"
)

func Transactions(c *gin.Context) {
	email, _ := c.Get("email")
	transactions, err := models.GetAllTransactions(email.(string))

	if err != nil {
		c.JSON(500, nil)
	}
	c.JSON(200, gin.H{
		"transactions": transactions,
	})
}
