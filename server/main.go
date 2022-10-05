package main

import (
	"fmt"
	"log"
	"os"
	handlers "server/handlers"
	"server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	APP_PORT string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}
	APP_PORT = os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "8000"
	}
}

func initRoutes() (r *gin.Engine) {
	r = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.GET("/test", handlers.Test)
		}
		private := api.Group("/private").Use(middleware.Authz())
		{
			private.GET("/test", handlers.Test)
			private.POST("/linktoken", handlers.CreateLinkToken)
			private.POST("/accesstoken", handlers.CreateAccessToken)
			private.GET("/gettransactions", handlers.GetPlaidTransactions)
			private.GET("/accounts", handlers.GetAccounts)
			private.POST("/removeaccount", handlers.RemoveAccount)
		}
	}
	return r
}

func main() {
	r := initRoutes()

	err := r.Run(":" + APP_PORT)
	if err != nil {
		log.Fatal(err)
	}
}
