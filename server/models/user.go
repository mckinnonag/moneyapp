package models

import (
	"database/sql"
	"errors"
)

var (
	DB *sql.DB
)

func SetAccessToken(uid, access_token, plaid_item_id string) error {
	sqlStatement := `
	INSERT INTO plaid_items (user_id, access_token, plaid_item_id)
	VALUES ($1, $2, $3)`
	_, err := DB.Exec(sqlStatement, uid, access_token, plaid_item_id)
	if err != nil {
		return err
	}
	return nil
}

func GetAccessTokens(uid string) (tokens []string, err error) {
	var accessToken string
	var result []string
	sqlStatement := `SELECT access_token FROM plaid_items WHERE user_id=$1`
	row := DB.QueryRow(sqlStatement, uid)
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
