package main

import (
	"database/sql"
	"fmt"
	"log"
	"server/common"
	handlers "server/handlers"
	"server/middleware"
	models "server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
			public.POST("/register", handlers.Register)
		}
		private := api.Group("/private").Use(middleware.Authz())
		{
			private.POST("/linktoken", handlers.CreateLinkToken)
			private.POST("/accesstoken", handlers.CreateAccessToken)
			private.GET("/gettransactions", handlers.GetTransactions)
			private.GET("/profile", handlers.Profile)
			private.GET("/accounts", handlers.GetAccounts)
			private.POST("/removeaccount", handlers.RemoveAccount)
		}
	}
	return r
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		common.DATABASE_URL, common.DATABASE_PORT, common.DATABASE_USER, common.DATABASE_PW, common.DATABASE_NAME)

	var err error
	models.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer models.DB.Close()

	err = models.DB.Ping()
	if err != nil {
		panic(err)
	}

	r := initRoutes()

	err = r.Run(":" + common.APP_PORT)
	if err != nil {
		panic("unable to start server")
	}
}
