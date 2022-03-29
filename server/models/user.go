package models

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
	{Username: "user3", Password: "pass3"},
}

func RegisterNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	u := user{Username: username, Password: string(hashedPassword)}

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
		// if u.Username == username && u.Password == password {
		// 	return true
		// }
		if u.Username == username && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
			return true
		}
	}
	return false
}
