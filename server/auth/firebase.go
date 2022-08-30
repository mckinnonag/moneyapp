package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func initializeAppDefault() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	app := initializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	// log.Printf("Verified ID token: %v\n", token)

	return token, nil
}

// Generate a custom token used for testing
func CreateCustomToken(ctx context.Context, uid string) (string, error) {
	app := initializeAppDefault()
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.CustomToken(ctx, uid)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Create a Firebase user
func CreateUser(ctx context.Context, email string, password string, displayName string) (*auth.UserRecord, error) {
	app := initializeAppDefault()
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		// PhoneNumber("+15555550100").
		Password(password).
		DisplayName(displayName).
		// PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func DeleteUser(ctx context.Context, uid string) error {
	app := initializeAppDefault()
	client, err := app.Auth(context.Background())
	if err != nil {
		return err
	}
	err = client.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

// Send API request to Firebase to exchange a custom token for an ID token. See https://firebase.google.com/docs/reference/rest/auth/
func ExchangeCustomTokenForIDToken(customToken string, apiKey string) (string, error) {
	type payload struct {
		Token             string `json:"token"`
		ReturnSecureToken string `json:"returnSecureToken"` // Always true
	}

	type response struct {
		IdToken      string `json:"idToken"`      // A Firebase Auth ID token generated from the provided custom token.
		RefreshToken string `json:"refreshToken"` // A Firebase Auth refresh token generated from the provided custom token.
		ExpiresIn    string `json:"expiresIn"`    // The number of seconds in which the ID token expires.
	}

	pay, err := json.Marshal(payload{Token: customToken, ReturnSecureToken: "true"})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		"https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key="+apiKey,
		"application/json",
		bytes.NewBuffer(pay),
	)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	var data response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	return data.IdToken, nil
}
