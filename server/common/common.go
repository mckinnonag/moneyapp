package common

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	plaid "github.com/plaid/plaid-go/plaid"
)

var (
	PLAID_CLIENT_ID     string
	PLAID_SECRET        string
	PLAID_ENV           string
	PLAID_PRODUCTS      string
	PLAID_COUNTRY_CODES string
	PLAID_REDIRECT_URI  string
	APP_PORT            string
	JWT_SECRET          string
	JWT_ISSUER          string
	DATABASE_URL        string
	DATABASE_PORT       int
	DATABASE_USER       string
	DATABASE_NAME       string
	DATABASE_PW         string
	DATABASE_SSL        string
	JWT_EXPIRY          int64
	PlaidClient         *plaid.APIClient = nil
)

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	DATABASE_URL = os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = "localhost"
	}
	DATABASE_PORT, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	DATABASE_USER = os.Getenv("DATABASE_USER")
	DATABASE_PW = os.Getenv("DATABASE_PW")
	DATABASE_NAME = os.Getenv("DATABASE_NAME")
	DATABASE_SSL = os.Getenv("DATABASE_SSL")
	if DATABASE_USER == "" || DATABASE_PW == "" || DATABASE_NAME == "" {
		log.Fatal("DATABASE_USER or DATABASE_PW is not set")
	}
	if DATABASE_SSL == "" {
		DATABASE_SSL = "disable"
	}

	JWT_SECRET = os.Getenv("JWT_SECRET")
	JWT_ISSUER = os.Getenv("JWT_ISSUER")
	exp, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRY"), 10, 64)
	JWT_EXPIRY = exp
	if JWT_SECRET == "" {
		JWT_SECRET = "verysecretkey"
	}
	if JWT_ISSUER == "" {
		JWT_ISSUER = "Issuer"
	}
	if JWT_EXPIRY == 0 {
		JWT_EXPIRY = 1
	}

	PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID")
	PLAID_SECRET = os.Getenv("PLAID_SECRET")

	if PLAID_CLIENT_ID == "" || PLAID_SECRET == "" {
		log.Fatal("Error: PLAID_SECRET or PLAID_CLIENT_ID is not set.")
	}

	PLAID_ENV = os.Getenv("PLAID_ENV")
	PLAID_PRODUCTS = os.Getenv("PLAID_PRODUCTS")
	PLAID_COUNTRY_CODES = os.Getenv("PLAID_COUNTRY_CODES")
	PLAID_REDIRECT_URI = os.Getenv("PLAID_REDIRECT_URI")
	APP_PORT = os.Getenv("APP_PORT")

	if PLAID_PRODUCTS == "" {
		PLAID_PRODUCTS = "transactions"
	}
	if PLAID_COUNTRY_CODES == "" {
		PLAID_COUNTRY_CODES = "US"
	}
	if PLAID_ENV == "" {
		PLAID_ENV = "sandbox"
	}
	if APP_PORT == "" {
		APP_PORT = "8000"
	}
	if PLAID_CLIENT_ID == "" {
		log.Fatal("PLAID_CLIENT_ID is not set.")
	}
	if PLAID_SECRET == "" {
		log.Fatal("PLAID_SECRET is not set.")
	}

	// create Plaid client
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", PLAID_CLIENT_ID)
	configuration.AddDefaultHeader("PLAID-SECRET", PLAID_SECRET)
	configuration.UseEnvironment(environments[PLAID_ENV])
	PlaidClient = plaid.NewAPIClient(configuration)
}
