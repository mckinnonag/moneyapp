package main

import (
	"server/common"
	handlers "server/handlers"
	"server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	common.Init()
}

func initRoutes() (r *gin.Engine) {
	r = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", handlers.Login)
		}
		private := api.Group("/private").Use(middleware.Authz())
		{
			private.POST("/linktoken", handlers.CreateLinkToken)
			private.POST("/accesstoken", handlers.CreateAccessToken)
			private.GET("/transactions", handlers.GetTransactions)
			private.GET("/profile", handlers.Profile)
		}
	}
	return r
}

func main() {
	r := initRoutes()

	err := r.Run(":" + common.APP_PORT)
	if err != nil {
		panic("unable to start server")
	}
}
