package handlers

import (
	models "server/models"

	"github.com/gin-gonic/gin"
)

func GetSharedTransactions(c *gin.Context) {
	uid, _ := c.Get("uid")
	transactions, err := models.GetSharedTransactions(uid.(string))

	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, gin.H{
		"transactions": transactions,
	})
}

func ShareTransaction(c *gin.Context) {

}
