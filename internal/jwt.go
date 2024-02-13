package internal

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTClaim struct {
	jwt.RegisteredClaims

	Name  string `json:"name"`
	Email string `json:"email"`
}

func generateJWT(email string, name string, getenv func(string) string) (tokenString string, err error) {
	if email == "" {
		return "", fmt.Errorf("email is required")
	}

	if name == "" {
		return "", fmt.Errorf("name is required")
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		Email: email,
		Name:  name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := getenv("SECRET")
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func validateToken(signedToken string, getenv func(string) string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, fmt.Errorf("couldn't parse claims")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
