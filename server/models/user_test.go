package models

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestUsernameAvailability(t *testing.T) {
	SaveLists()

	if !isUsernameAvailable("newuser") {
		t.Fail()
	}

	if isUsernameAvailable("user1") {
		t.Fail()
	}

	registerNewUser("newuser", "newpass")

	if isUsernameAvailable("newuser") {
		t.Fail()
	}

	RestoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	SaveLists()

	u, err := registerNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	RestoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	SaveLists()

	u, err := registerNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = registerNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	RestoreLists()
}
