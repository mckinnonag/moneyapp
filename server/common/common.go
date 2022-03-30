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
	PLAID_CLIENT_ID     = ""
	PLAID_SECRET        = ""
	PLAID_ENV           = ""
	PLAID_PRODUCTS      = ""
	PLAID_COUNTRY_CODES = ""
	PLAID_REDIRECT_URI  = ""
	APP_PORT            = ""
	// TEMPLATE_PATH       = ""
	JWT_SECRET  = ""
	JWT_ISSUER  = ""
	JWT_EXPIRY  int64
	PlaidClient *plaid.APIClient = nil
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

	JWT_SECRET = os.Getenv("JWT_SECRET")
	JWT_ISSUER = os.Getenv("JWT_ISSUER")
	exp, err := strconv.ParseInt(os.Getenv("JWT_EXPIRY"), 10, 64)
	if err != nil {
		log.Println(err)
	}
	JWT_EXPIRY = exp

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

	// if TEMPLATE_PATH == "" {
	// 	TEMPLATE_PATH = "/Users/alexmckinnon/moneytest/templates/*"
	// }
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

// func Render(c *gin.Context, data gin.H, templateName string) {
// 	switch c.Request.Header.Get("Accept") {
// 	case "application/json":
// 		// Respond with JSON
// 		c.JSON(http.StatusOK, data["payload"])
// 	case "application/xml":
// 		// Respond with XML
// 		c.XML(http.StatusOK, data["payload"])
// 	default:
// 		// Respond with HTML
// 		c.HTML(http.StatusOK, templateName, data)
// 	}

// }
