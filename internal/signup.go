package internal

import (
	"github.com/go-fuego/fuego"
	"golang.org/x/crypto/bcrypt"
	"reportpipe/internal/auth"
)

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (r *Routes) signup(c *fuego.ContextWithBody[SignUpRequest]) (TokenResponse, error) {
	req, err := c.Body()
	if err != nil {
		return TokenResponse{}, err
	}

	pw, err := hashPassword(req.Password)
	if err != nil {
		return TokenResponse{}, err
	}

	_, err = r.DB.ExecContext(c.Context(), "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", req.Username, req.Email, pw)
	if err != nil {
		return TokenResponse{}, err
	}

	token, err := auth.GenerateJWT(req.Email, req.Username, r.getEnv)
	if err != nil {
		return TokenResponse{}, err
	}

	c.SetStatus(201)
	return TokenResponse{Token: token}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
