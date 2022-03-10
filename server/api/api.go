package api

import (
	handlers "server/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/", handlers.ShowIndexPage)
	userRoutes := r.Group("/u")
	{
		userRoutes.GET("/register", handlers.ShowRegistrationPage)
		userRoutes.POST("/register", handlers.Register)
		userRoutes.GET("/login", handlers.ShowLoginPage)
		userRoutes.POST("/login", handlers.PerformLogin)
		userRoutes.GET("/logout", handlers.Logout)
	}

}
