package main

import (
	"os"
	api "server/api"
	"server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Initialize gin router and establish routes
func (s *server) routes() error {
	APP_PORT = os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "8000"
	}

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "")
	})

	handlers := r.Group("/api")
	{
		public := handlers.Group("/public")
		{
			public.GET("/test", api.Test)
		}
		private := handlers.Group("/private").Use(middleware.Authz())
		{
			private.GET("/test", api.Test)
			private.POST("/linktoken", api.CreateLinkToken)
			private.POST("/accesstoken", api.CreateAccessToken)
			private.GET("/gettransactions", api.GetPlaidTransactions)
			private.GET("/getsharedtransactions", api.GetTransactions)
			private.GET("/accounts", api.GetAccounts)
			private.POST("/removeaccount", api.RemoveAccount)
		}
	}
	s.router = r
	err := r.Run(":" + APP_PORT)
	if err != nil {
		return err
	}
	return nil
}
