package handlers

import (
	"encoding/json"
	models "server/models"

	"github.com/gin-gonic/gin"
)

func Transactions(c *gin.Context) {
	email, _ := c.Get("email")
	transactions := models.GetAllTransactions(email.(string))
	data, _ := json.Marshal(&transactions)

	c.Data(200, "application/json", data)
}
