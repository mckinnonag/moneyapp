package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"moneyapp/pkg/api"
	"moneyapp/pkg/app"
	"moneyapp/pkg/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	APP_PORT      string
	DATABASE_URL  string
	DATABASE_PORT int
	DATABASE_USER string
	DATABASE_NAME string
	DATABASE_PW   string
	DATABASE_SSL  string
)

func dbConnectionString() (string, error) {
	DATABASE_URL = os.Getenv("POSTGRES_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = "localhost"
	}
	DATABASE_PORT, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	DATABASE_USER = os.Getenv("POSTGRES_USER")
	DATABASE_PW = os.Getenv("POSTGRES_PASSWORD")
	DATABASE_NAME = os.Getenv("POSTGRES_NAME")
	DATABASE_SSL = os.Getenv("POSTGRES_SSL")
	if DATABASE_USER == "" || DATABASE_PW == "" || DATABASE_NAME == "" {
		return "", errors.New("DATABASE_USER or DATABASE_PW is not set")
	}
	if DATABASE_SSL == "" {
		DATABASE_SSL = "disable"
	}

	// Initialize DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		DATABASE_URL, DATABASE_PORT, DATABASE_USER, DATABASE_PW, DATABASE_NAME, DATABASE_SSL)
	return psqlInfo, nil
}

// Initialize the DB connection
func connectDB(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Connect DB
	connectionString, err := dbConnectionString()
	if err != nil {
		return err
	}
	db, err := connectDB(connectionString)
	if err != nil {
		return err
	}

	// Create storage dependency
	storage := repository.NewStorage(db)

	// Run database migrations
	err = storage.RunMigrations(connectionString)

	if err != nil {
		return err
	}

	// Create router dependency
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

	transactionService := api.NewTransactionService(storage)

	plaidConfig := &api.PlaidConfig{
		APP_NAME:            os.Getenv("APP_NAME"),
		PLAID_CLIENT_ID:     os.Getenv("PLAID_CLIENT_ID"),
		PLAID_SECRET:        os.Getenv("PLAID_SECRET"),
		PLAID_ENV:           os.Getenv("PLAID_ENV"),
		PLAID_PRODUCTS:      os.Getenv("PLAID_PRODUCTS"),
		PLAID_COUNTRY_CODES: os.Getenv("PLAID_COUNTRY_CODES"),
		PLAID_REDIRECT_URI:  os.Getenv("PLAID_REDIRECT_URI"),
	}
	plaidService := api.NewPlaidService(plaidConfig, storage)

	server := app.NewServer(r, transactionService, plaidService)

	// Start the server
	err = server.Run(APP_PORT)
	if err != nil {
		return err
	}

	return nil
}
