package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-fuego/fuego"

	"github.com/golang-jwt/jwt/v5"
)

var _ jwt.Claims = &JWTClaim{}

type JWTClaim struct {
	jwt.RegisteredClaims

	Username string `json:"username"`
	Email    string `json:"email"`
	// Roles    []string  // TODO: Add roles to the token.
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

type Auth struct {
	DB     *sql.DB
	GetEnv func(string) string
}

// Middleware validates the token and calls the next handler.
func (a *Auth) Middleware(next http.Handler) http.Handler {
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

		secret := a.GetEnv("SECRET")
		claims, err := validateToken(authHeader[1], func(string) string {
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

		resp, err := a.DB.QueryContext(r.Context(), "SELECT revoke_token_before FROM users WHERE username = $1", claims.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fuego.SendJSON(w, map[string]string{"error": "Internal Server Error"})
			return
		}

		defer resp.Close()

		if !resp.Next() {
			w.WriteHeader(http.StatusUnauthorized)
			fuego.SendJSON(w, map[string]string{"error": "User not found"})
			return
		}

		var revokeTokenBefore time.Time
		err = resp.Scan(&revokeTokenBefore)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusInternalServerError)
			fuego.SendJSON(w, map[string]string{"error": err.Error()})
			return
		}

		if claims.IssuedAt.Time.Before(revokeTokenBefore) {
			w.WriteHeader(http.StatusUnauthorized)
			fuego.SendJSON(w, map[string]string{"error": "Token revoked"})
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "email", claims.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
