package api

import (
	"context"
	"errors"
	"log"
	"strconv"
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

// PlaidService contains the methods of the Plaid service
type PlaidService interface {
	CreateLinkToken(c *gin.Context) (string, error)
	GetAccessToken(c *gin.Context, a *NewAccessTokenRequest) (string, string, error)
	GetPlaidTransactions(c *gin.Context, req GetPlaidTransactionsRequest) ([]plaid.Transaction, error)
}

// PlaidRepository is what lets our service do db operations without knowing anything about the implementation
type PlaidRepository interface {
	CreateAccessToken(*NewAccessTokenRequest) error
	GetAccessTokens(string) ([]string, error)
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

func (p *plaidService) CreateLinkToken(c *gin.Context) (string, error) {
	ctx := context.Background()
	countryCodes := convertCountryCodes(strings.Split(plaid_country_codes, ","))
	products := convertProducts(strings.Split(plaid_products, ","))
	redirectURI := plaid_redirect_uri

	uid, exists := c.Get("uid")
	if !exists {
		return "", errors.New("request context does not contain user id claim")
	}

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: uid.(string),
	}

	request := plaid.NewLinkTokenCreateRequest(
		app_name,
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
		return "", err
	}

	return linkTokenCreateResp.GetLinkToken(), nil
}

func (p *plaidService) GetAccessToken(c *gin.Context, a *NewAccessTokenRequest) (accessToken, itemID string, err error) {
	publicToken := a.PublicToken
	if publicToken == "" {
		err = errors.New("missing public token")
		return
	}
	ctx := context.Background()

	// exchange the public_token for an access_token
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()
	if err != nil {
		return
	}

	a.AccessToken = exchangePublicTokenResp.GetAccessToken()
	a.ItemId = exchangePublicTokenResp.GetItemId()

	err = p.storage.CreateAccessToken(a)
	if err != nil {
		return
	}

	accessToken = a.AccessToken
	itemID = a.ItemId
	return
}

// Generate a public token for testing (sandbox only). TODO: do not export this
func CreatePublicToken() (string, error) {
	if plaid_env != "sandbox" {
		return "", errors.New("can only generate a public token for testing")
	}

	ctx := context.Background()
	products := convertProducts(strings.Split(plaid_products, ","))
	institution := "ins_109508"

	// Get a random public token
	sandboxPublicTokenResp, _, err := client.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
		*plaid.NewSandboxPublicTokenCreateRequest(
			institution,
			products,
		),
	).Execute()

	if err != nil {
		return "", err
	}

	return sandboxPublicTokenResp.GetPublicToken(), nil
}

func (p *plaidService) GetPlaidTransactions(c *gin.Context, req GetPlaidTransactionsRequest) ([]plaid.Transaction, error) {
	ctx := context.Background()

	accessTokens, err := p.storage.GetAccessTokens(req.UID)
	if err != nil {
		return nil, err
	}

	var ret []plaid.Transaction
	for _, accessToken := range accessTokens {
		request := plaid.NewTransactionsGetRequest(
			accessToken,
			req.StartDate,
			req.EndDate,
		)

		count, err := strconv.Atoi(req.Count)
		if err != nil {
			return nil, err
		}
		offset, err := strconv.Atoi(req.Offset)
		if err != nil {
			return nil, err
		}
		options := plaid.TransactionsGetRequestOptions{
			Count:  plaid.PtrInt32(int32(count)),
			Offset: plaid.PtrInt32(int32(offset)),
		}

		request.SetOptions(options)

		res, _, err := client.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(*request).Execute()

		if err != nil {
			return nil, err
		}

		ret = append(ret, res.Transactions...)
	}
	return ret, nil
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

// Helper function to extract the error message body
func plaidError(e error) error {
	plaidErr, err := plaid.ToPlaidError(e)
	if e != nil {
		return err
	}
	return errors.New(plaidErr.ErrorMessage)
}
