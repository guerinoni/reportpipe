package internal

import (
	"github.com/go-fuego/fuego"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"reportpipe/internal/auth"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *Routes) login(c *fuego.ContextWithBody[LoginRequest]) (TokenResponse, error) {
	req, err := c.Body()
	if err != nil {
		return TokenResponse{}, err
	}

	rows, err := r.DB.QueryContext(c.Context(), "SELECT username,password FROM users WHERE email = $1", req.Email)
	if err != nil {
		return TokenResponse{}, fuego.ErrUnauthorized
	}

	defer rows.Close()

	if !rows.Next() {
		c.SetStatus(401)
		return TokenResponse{}, fuego.ErrUnauthorized
	}

	var username, pw string
	if err := rows.Scan(&username, &pw); err != nil {
		return TokenResponse{}, fuego.ErrUnauthorized
	}

	if err := checkPassword(req.Password, pw); err != nil {
		return TokenResponse{}, fuego.ErrUnauthorized
	}

	token, err := auth.GenerateJWT(req.Email, username, r.getEnv)
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{Token: token}, err
}

func checkPassword(providedPassword string, pw string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
