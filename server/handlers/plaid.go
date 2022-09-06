package handlers

import (
	"context"
	"log"
	"net/http"
	common "server/common"
	"server/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	plaid "github.com/plaid/plaid-go/plaid"
)

func renderError(c *gin.Context, originalErr error) {
	if plaidError, err := plaid.ToPlaidError(originalErr); err == nil {
		// Return 200 and allow the front end to render the error.
		c.JSON(http.StatusInternalServerError, gin.H{"error": plaidError})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": originalErr.Error()})
}

func convertCountryCodes(countryCodeStrs []string) []plaid.CountryCode {
	countryCodes := []plaid.CountryCode{}

	for _, countryCodeStr := range countryCodeStrs {
		countryCodes = append(countryCodes, plaid.CountryCode(countryCodeStr))
	}

	return countryCodes
}

func convertProducts(productStrs []string) []plaid.Products {
	products := []plaid.Products{}

	for _, productStr := range productStrs {
		products = append(products, plaid.Products(productStr))
	}

	return products
}

func linkTokenCreate(
	paymentInitiation *plaid.LinkTokenCreateRequestPaymentInitiation,
) (string, error) {
	ctx := context.Background()
	countryCodes := convertCountryCodes(strings.Split(common.PLAID_COUNTRY_CODES, ","))
	products := convertProducts(strings.Split(common.PLAID_PRODUCTS, ","))
	redirectURI := common.PLAID_REDIRECT_URI

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: time.Now().String(),
	}

	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Quickstart",
		"en",
		countryCodes,
		user,
	)

	request.SetProducts(products)

	if redirectURI != "" {
		request.SetRedirectUri(redirectURI)
	}

	if paymentInitiation != nil {
		request.SetPaymentInitiation(*paymentInitiation)
	}

	linkTokenCreateResp, _, err := common.PlaidClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	if err != nil {
		return "", err
	}

	return linkTokenCreateResp.GetLinkToken(), nil
}

func getAccessToken(c *gin.Context, publicToken string) {
	email, _ := c.Get("email")
	ctx := context.Background()

	// exchange the public_token for an access_token
	exchangePublicTokenResp, _, err := common.PlaidClient.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()
	if err != nil {
		renderError(c, err)
		return
	}

	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemID := exchangePublicTokenResp.GetItemId()

	err = models.SetAccessToken(email.(string), accessToken, itemID)
	if err != nil {
		renderError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"item_id":      itemID,
	})
}

func CreateLinkToken(c *gin.Context) {
	linkToken, err := linkTokenCreate(nil)
	if err != nil {
		renderError(c, err)
		return
	}
	c.JSON(200, gin.H{"link_token": linkToken})
}

func CreateAccessToken(c *gin.Context) {
	type TOKEN struct {
		Token string `json:"token" binding: "required"`
	}
	var token TOKEN
	err := c.BindJSON(&token)
	if err != nil {
		renderError(c, err)
	}

	getAccessToken(c, token.Token)
	c.JSON(http.StatusOK, nil)
}

// Helper function to determine if Transfer is in Plaid product array
// func itemExists(array []string, product string) bool {
// 	for _, item := range array {
// 		if item == product {
// 			return true
// 		}
// 	}

// 	return false
// }

func GetAccounts(c *gin.Context) {
	type accountInfo struct {
		AccountId        string
		BalanceAvailable float32
		Name             string
		OfficialName     string
		Subtype          plaid.NullableAccountSubtype
		Type             string
	}

	// Call the Plaid API to get a list of accounts
	ctx := context.Background()
	email, exists := c.Get("email")
	if !exists {
		log.Println("email doesn't exist")
	}
	accessTokens, err := models.GetAccessTokens(email.(string))
	if err != nil {
		renderError(c, err)
	}

	var accts []accountInfo

	for _, accessToken := range accessTokens {
		accountsGetRequest := plaid.NewAccountsGetRequest(accessToken)
		// accountsGetRequest.SetOptions(plaid.AccountsGetRequestOptions{
		// 	AccountIds: &[]string{},
		// })
		accountsGetResp, _, _ := common.PlaidClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
			*accountsGetRequest,
		).Execute()
		response := accountsGetResp.GetAccounts()

		for _, a := range response {
			accountId := a.AccountId
			var balanceAvailable float32
			if a.Balances.Available.Get() != nil {
				balanceAvailable = *a.Balances.Available.Get()
			}

			var officialName string

			if a.OfficialName.Get() != nil {
				officialName = *a.OfficialName.Get()
			}

			acct := accountInfo{
				AccountId:        accountId,
				BalanceAvailable: balanceAvailable,
				Name:             a.Name,
				OfficialName:     officialName,
				// Subtype:          *a.Subtype.Get(),
				// Type:             *a.Type.Get(),
			}
			accts = append(accts, acct)
		}
	}

	c.JSON(200, gin.H{
		"accounts": accts,
	})
}

func RemoveAccount(c *gin.Context) {
	ctx := context.Background()
	email, exists := c.Get("email")
	if !exists {
		log.Println("email doesn't exist")
	}
	accessTokens, err := models.GetAccessTokens(email.(string))
	if err != nil {
		renderError(c, err)
	}

	for _, accessToken := range accessTokens {
		request := plaid.NewItemRemoveRequest(accessToken)
		_, _, err = common.PlaidClient.PlaidApi.ItemRemove(ctx).ItemRemoveRequest(*request).Execute()
		if err != nil {
			log.Println(err)
		}
	}
}

func GetTransactions(c *gin.Context) {
	const iso8601TimeFormat = "2006-01-02"
	startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDate := time.Now().Format(iso8601TimeFormat)

	ctx := context.Background()
	email, exists := c.Get("email")
	if !exists {
		log.Println("email doesn't exist")
	}
	accessTokens, err := models.GetAccessTokens(email.(string))
	if err != nil {
		c.JSON(500, nil)
		return
	}

	var transactions []models.Transaction

	for _, accessToken := range accessTokens {
		request := plaid.NewTransactionsGetRequest(
			accessToken,
			startDate,
			endDate,
		)

		options := plaid.TransactionsGetRequestOptions{
			Count:  plaid.PtrInt32(100),
			Offset: plaid.PtrInt32(0),
		}

		request.SetOptions(options)

		res, _, err := common.PlaidClient.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(*request).Execute()

		if err != nil {
			c.JSON(500, nil)
			return
		}

		for _, t := range res.Transactions {
			var merchant string
			if t.MerchantName.Get() != nil {
				merchant = *t.MerchantName.Get()
			}

			tx := models.Transaction{
				ID:           t.TransactionId,
				MerchantName: merchant,
				Amount:       t.Amount,
				Date:         t.Date,
			}
			transactions = append(transactions, tx)
		}
	}

	c.JSON(200, gin.H{
		"transactions": transactions,
	})
}
