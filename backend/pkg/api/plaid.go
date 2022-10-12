package api

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	plaid "github.com/plaid/plaid-go/plaid"
)

var (
	app_name            string
	plaid_client_id     string
	plaid_secret        string
	plaid_env           string
	plaid_products      string
	plaid_country_codes string
	plaid_redirect_uri  string
	client              *plaid.APIClient = nil
)

type PlaidConfig struct {
	APP_NAME            string
	PLAID_CLIENT_ID     string
	PLAID_SECRET        string
	PLAID_ENV           string
	PLAID_PRODUCTS      string
	PLAID_COUNTRY_CODES string
	PLAID_REDIRECT_URI  string
}

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

// Initializes global variables and plaid client
func start(config *PlaidConfig) {
	app_name = config.APP_NAME
	if app_name == "" {
		app_name = "plaid test"
	}

	plaid_client_id = config.PLAID_CLIENT_ID
	plaid_secret = config.PLAID_SECRET

	if plaid_client_id == "" || plaid_secret == "" {
		log.Fatal("Error: plaid_secret or plaid_client_id is not set.")
	}

	plaid_env = config.PLAID_ENV
	plaid_products = config.PLAID_PRODUCTS
	plaid_country_codes = config.PLAID_COUNTRY_CODES
	plaid_redirect_uri = config.PLAID_REDIRECT_URI

	if plaid_products == "" {
		plaid_products = "transactions"
	}
	if plaid_country_codes == "" {
		plaid_country_codes = "US"
	}
	if plaid_env == "" {
		plaid_env = "sandbox"
	}
	if plaid_client_id == "" {
		log.Fatal("plaid_client_id is not set.")
	}
	if plaid_secret == "" {
		log.Fatal("plaid_secret is not set.")
	}

	// create Plaid client
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", plaid_client_id)
	configuration.AddDefaultHeader("PLAID-SECRET", plaid_secret)
	configuration.UseEnvironment(environments[plaid_env])
	client = plaid.NewAPIClient(configuration)
}

// PlaidService contains the methods of the Plaid service
type PlaidService interface {
	CreateLinkToken(c *gin.Context) (string, error)
}

// PlaidRepository is what lets our service do db operations without knowing anything about the implementation
type PlaidRepository interface {
	// CreatePlaid(NewPlaidRequest) error
}

type plaidService struct {
	storage PlaidRepository
}

func NewPlaidService(config *PlaidConfig, plaidRepo PlaidRepository) PlaidService {
	start(config)
	return &plaidService{
		storage: plaidRepo,
	}
}

func (p *plaidService) CreateLinkToken(c *gin.Context) (string, error) {
	ctx := context.Background()
	countryCodes := convertCountryCodes(strings.Split(plaid_country_codes, ","))
	products := convertProducts(strings.Split(plaid_products, ","))
	redirectURI := plaid_redirect_uri

	// uid, exists := c.Get("uid")
	// if !exists {
	// 	return "", errors.New("request context does not contain user id claim")
	// }

	// user := plaid.LinkTokenCreateRequestUser{
	// 	ClientUserId: uid.(string),
	// }

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: "testuid123456",
	}

	request := plaid.NewLinkTokenCreateRequest(
		"Moneyapp",
		"en",
		countryCodes,
		user,
	)

	request.SetProducts(products)
	if redirectURI != "" {
		request.SetRedirectUri(redirectURI)
	}

	linkTokenCreateResp, _, err := client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	if err != nil {
		plaidErr, e := plaid.ToPlaidError(err)
		if e != nil {
			return "", e
		}
		return "", errors.New(plaidErr.ErrorMessage)
	}

	return linkTokenCreateResp.GetLinkToken(), nil
}

// Helper function to convert string to plaid country code
func convertCountryCodes(countryCodeStrs []string) []plaid.CountryCode {
	countryCodes := []plaid.CountryCode{}

	for _, countryCodeStr := range countryCodeStrs {
		countryCodes = append(countryCodes, plaid.CountryCode(countryCodeStr))
	}

	return countryCodes
}

// Helper function to convert product string to plaid product
func convertProducts(productStrs []string) []plaid.Products {
	products := []plaid.Products{}

	for _, productStr := range productStrs {
		products = append(products, plaid.Products(productStr))
	}

	return products
}
