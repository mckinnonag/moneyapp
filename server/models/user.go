package models

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func testHash(pass string) []byte {
	result, _ := bcrypt.GenerateFromPassword([]byte("pass1"), 8)
	return result
}

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []User{
	{Username: "user1@user.com", Password: string(testHash("pass1"))},
	{Username: "user1@user.com", Password: string(testHash("pass2"))},
	{Username: "user1@user.com", Password: string(testHash("pass3"))},
}

func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	u := User{Username: username, Password: string(hashedPassword)}

	userList = append(userList, u)

	return &u, nil
}

func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}

func IsUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
			return true
		}
	}
	return false
}
