package api

import (
	auth "server/auth"
	handlers "server/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Use(auth.SetUserStatus())

	r.GET("/", handlers.ShowIndexPage)

	// Testing
	r.GET("/login", handlers.ReactLogin) // Redundant with /u/ route below

	// type LOGIN struct {
	// 	USER     string `json:"username"` //binding:"required"`
	// 	PASSWORD string `json:"password"` //binding:"required"`
	// }
	// r.POST("/login", func(c *gin.Context) {
	// 	var credentials LOGIN
	// 	c.BindJSON((&credentials))
	// 	c.JSON(
	// 		200,
	// 		gin.H{"status": credentials.USER},
	// 	)
	// })
	r.POST("/login", handlers.ReactPerformLogin)

	userRoutes := r.Group("/u")
	{
		userRoutes.GET("/register", auth.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)
		userRoutes.POST("/register", auth.EnsureNotLoggedIn(), handlers.Register)
		userRoutes.GET("/login", auth.EnsureNotLoggedIn(), handlers.ShowLoginPage)
		userRoutes.POST("/login", auth.EnsureNotLoggedIn(), handlers.PerformLogin)
		userRoutes.GET("/logout", auth.EnsureLoggedIn(), handlers.Logout)
	}

}
