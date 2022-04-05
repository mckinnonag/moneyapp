package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	DB *sql.DB
)

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
	INSERT INTO users (email, password, createdon)
	VALUES ($1, $2, $3)`
	_, err = DB.Exec(sqlStatement, username, hashedPassword, time.Now())
	if err != nil {
		panic(err)
	}

	return nil
}

func isUsernameAvailable(email string) bool {
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
