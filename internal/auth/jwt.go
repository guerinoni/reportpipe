package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var _ jwt.Claims = &JWTClaim{}

type JWTClaim struct {
	jwt.RegisteredClaims

	Username string `json:"username"`
	Email    string `json:"email"`
	//Roles    []string  // TODO: Add roles to the token.
}

// GenerateJWT creates a new JWT token.
func GenerateJWT(email string, username string, getenv func(string) string) (tokenString string, err error) {
	if email == "" {
		return "", fmt.Errorf("email is required")
	}
	if username == "" {
		return "", fmt.Errorf("username is required")
	}

	claims := &JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "reportpipe",
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Email:    email,
		Username: username,
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

// Middleware validates the token and calls the next handler.
func Middleware(next http.Handler, getenv func(string) string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Malformed Token"))
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		secret := getenv("SECRET")
		_, err := validateToken(authHeader[1], func(string) string {
			return secret
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Invalid Token"))
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}
