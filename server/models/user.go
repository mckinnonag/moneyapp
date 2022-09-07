package models

import (
	"database/sql"
	"errors"
)

var (
	DB *sql.DB
)

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
