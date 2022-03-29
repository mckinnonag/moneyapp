package api

import (
	auth "server/auth"
	handlers "server/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Use(auth.SetUserStatus())

	r.GET("/", handlers.ShowIndexPage)
	// r.GET("/login", handlers.Login)
	r.POST("/login", handlers.PerformLogin)

	api := r.Group("/api")
	{
		api.POST("/linktoken", handlers.CreateLinkToken)
		api.POST("/accesstoken", handlers.CreateAccessToken)
	}

	// userRoutes := r.Group("/u")
	// {
	// 	userRoutes.GET("/register", auth.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)
	// 	userRoutes.POST("/register", auth.EnsureNotLoggedIn(), handlers.Register)
	// 	userRoutes.GET("/login", auth.EnsureNotLoggedIn(), handlers.ShowLoginPage)
	// 	userRoutes.POST("/login", auth.EnsureNotLoggedIn(), handlers.PerformLogin)
	// 	userRoutes.GET("/logout", auth.EnsureLoggedIn(), handlers.Logout)
	// }
}
