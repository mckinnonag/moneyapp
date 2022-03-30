package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtConfig struct {
	SecretKey  string
	Issuer     string
	Expiration int64
}

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

func (j *JwtConfig) GenerateToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.Expiration)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	signedToken, err = token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return
	}
	return
}

func (j *JwtConfig) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return
	}

	return
}
