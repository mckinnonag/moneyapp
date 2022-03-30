package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	common "server/common"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	plaid "github.com/plaid/plaid-go/plaid"
)

// We store the access_token in memory - in production, store it in a secure
// persistent data store.
var accessToken string
var itemID string

// The transfer_id is only relevant for the Transfer ACH product.
// We store the transfer_id in memory - in production, store it in a secure
// persistent data store
var transferID string

// Temporary testing bucket for access tokens
var accessTokens map[string]string

type accountInfo struct {
	AccountId        string
	BalanceAvailable float32
	Name             string
	OfficialName     string
	Subtype          plaid.NullableAccountSubtype
	Type             string
}

func renderError(c *gin.Context, originalErr error) {
	if plaidError, err := plaid.ToPlaidError(originalErr); err == nil {
		// Return 200 and allow the front end to render the error.
		c.JSON(http.StatusOK, gin.H{"error": plaidError})
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
	// publicToken := c.PostForm("public_token")
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

	accessToken = exchangePublicTokenResp.GetAccessToken()
	itemID = exchangePublicTokenResp.GetItemId()
	// if itemExists(strings.Split(common.PLAID_PRODUCTS, ","), "transfer") {
	// 	transferID, err = authorizeAndCreateTransfer(ctx, common.PlaidClient, accessToken)
	// }

	fmt.Println("public token: " + publicToken)
	fmt.Println("access token: " + accessToken)
	fmt.Println("item ID: " + itemID)

	if accessTokens == nil {
		accessTokens = make(map[string]string)
	}
	accessTokens[email.(string)] = accessToken

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

func GetTransactions(c *gin.Context) {
	if accessToken == "" {
		return
	}

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

// type GetAccountsPayload struct {
// 	Email string `json:"email"`
// }

func GetAccounts(c *gin.Context) {
	// Call the Plaid API to get a list of accounts
	ctx := context.Background()
	email, exists := c.Get("email")
	if !exists {
		log.Println("email doesn't exist")
	}
	accessToken = accessTokens[email.(string)]
	accessToken = "access-sandbox-410dea3e-79b1-4c2a-8935-1164190fc552"
	fmt.Println(accessToken)
	accountsGetRequest := plaid.NewAccountsGetRequest(accessToken)
	// accountsGetRequest.SetOptions(plaid.AccountsGetRequestOptions{
	// 	AccountIds: &[]string{},
	// })
	accountsGetResp, _, _ := common.PlaidClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*accountsGetRequest,
	).Execute()
	response := accountsGetResp.GetAccounts()

	var accts []accountInfo
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

	// fmt.Println(accts)

	// body, err := json.Marshal(accts[0])
	// if err != nil {
	// 	fmt.Println(err)
	// }

	c.JSON(200, gin.H{
		"accounts": accts,
	})
}
