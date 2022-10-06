package api

import (
	models "server/models"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	uid, _ := c.Get("uid")
	transactions, err := models.GetTransactions(uid.(string))

	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, gin.H{
		"transactions": transactions,
	})
}

func CreateTransaction(c *gin.Context) {

}
