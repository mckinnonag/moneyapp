package auth

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenToken(t *testing.T) {
	j := JwtConfig{
		SecretKey:  "youllneverguess",
		Issuer:     "Service",
		Expiration: 2,
	}

	tok, err := j.GenerateToken()
	assert.NoError(t, err)

	os.Setenv("testTok", tok)
}

func TestValToken(t *testing.T) {
	tok := os.Getenv("testTok")

	j := JwtConfig{
		SecretKey: "youllneverguess",
		Issuer:    "Service",
	}

	err := j.ValidateToken(tok)
	assert.NoError(t, err)
}
