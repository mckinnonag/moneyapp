package handlers

import (
	"context"
	"fmt"
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
	c.JSON(http.StatusOK, gin.H{"link_token": linkToken})
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
