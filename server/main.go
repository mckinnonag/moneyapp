package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"server/common"
	handlers "server/handlers"
	"server/middleware"
	models "server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		c.String(200, "")
	})

	api := r.Group("/api")
	{
		// public := api.Group("/public")
		{
			// public.POST("/login", handlers.Login)
			// public.POST("/register", handlers.Register)
			// public.POST("/test", handlers.Test)
		}
		private := api.Group("/private").Use(middleware.Authz())
		// private := api.Group("/private")
		{
			private.POST("/test", handlers.Test)
			private.POST("/linktoken", handlers.CreateLinkToken)
			private.POST("/accesstoken", handlers.CreateAccessToken)
			private.GET("/gettransactions", handlers.GetTransactions)
			private.GET("/accounts", handlers.GetAccounts)
			private.POST("/removeaccount", handlers.RemoveAccount)
		}
	}
	return r
}

func main() {
	// DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		common.DATABASE_URL, common.DATABASE_PORT, common.DATABASE_USER, common.DATABASE_PW, common.DATABASE_NAME)

	var err error
	models.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer models.DB.Close()

	if err = models.DB.Ping(); err != nil {
		log.Fatal("unable to ping database")
	}

	// DB Migration
	var migrationDir = flag.String("migration.files", "../db/migration", "Directory where the migration files are located ?")
	driver, err := postgres.WithInstance(models.DB, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	file := "000001_init_schema.up.sql://../db/migration"
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir),
		file, driver)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	log.Println("Database Migrated!")

	// Handlers
	r := initRoutes()

	err = r.Run(":" + common.APP_PORT)
	if err != nil {
		panic(err)
	}
}
