package handlers

import (
	common "server/common"
	models "server/models"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	transactions := models.GetAllTransactions()

	common.Render(c, gin.H{
		"title":   "Home Page",
		"payload": transactions,
	},
		"index.html")
}
