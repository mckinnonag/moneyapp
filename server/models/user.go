package models

import (
	"database/sql"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	DB *sql.DB
)

func isUsernameAvailable(email string) bool {
	// Helper function to check if username is available
	sqlStatement := `SELECT email FROM users WHERE email=$1;`
	row := DB.QueryRow(sqlStatement, email)
	switch err := row.Scan(&email); err {
	case sql.ErrNoRows:
		return true
	case nil:
		return false
	default:
		panic(err)
	}
}

func lookupUser(email string) (user_id string, err error) {
	// Helper function. Accepts email; returns user_id
	sqlStatement := `SELECT user_id, email FROM users WHERE email=$1;`
	row := DB.QueryRow(sqlStatement, email)
	switch err = row.Scan(&user_id, &email); err {
	case sql.ErrNoRows:
		return "", errors.New("user does not exist")
	case nil:
		return user_id, nil
	default:
		return "", err
	}
}

func RegisterNewUser(username, password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("the password can't be empty")
	} else if !isUsernameAvailable(username) {
		return errors.New("the username isn't available")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return errors.New(err.Error())
	}

	sqlStatement := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)`
	_, err = DB.Exec(sqlStatement, username, hashedPassword)
	if err != nil {
		panic(err)
	}

	return nil
}

func IsUserValid(email string, password string) bool {
	sqlStatement := `SELECT password, email FROM users WHERE email=$1;`
	var pwHash string
	row := DB.QueryRow(sqlStatement, email)
	switch err := row.Scan(&pwHash, &email); err {
	case sql.ErrNoRows: // User does not exist
		return false
	case nil:
		match := bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(password))
		if match == nil {
			return true
		} else {
			return false
		}
	default:
		panic(err)
	}
}

func SetAccessToken(email, access_token, plaid_item_id string) error {
	user_id, err := lookupUser(email)
	if err != nil {
		return err
	}

	sqlStatement := `
	INSERT INTO plaid_items (user_id, access_token, plaid_item_id)
	VALUES ($1, $2, $3)`
	_, err = DB.Exec(sqlStatement, user_id, access_token, plaid_item_id)
	if err != nil {
		return err
	}
	return nil
}

func GetAccessTokens(email string) (tokens []string, err error) {
	user_id, err := lookupUser(email)
	if err != nil {
		return nil, err
	}

	var accessToken string
	var result []string
	sqlStatement := `SELECT access_token FROM plaid_items WHERE user_id=$1`
	row := DB.QueryRow(sqlStatement, user_id)
	switch err := row.Scan(&accessToken); err {
	case sql.ErrNoRows:
		return nil, errors.New("no results")
	case nil:
		result = append(result, accessToken)
	default:
		return nil, err
	}
	return result, nil
}
