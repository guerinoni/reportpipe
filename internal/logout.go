package internal

import "github.com/go-fuego/fuego"

type LogoutRequest struct {
	Token string `json:"token" validate:"required"`
}

func (r *Routes) logout(c *fuego.ContextWithBody[LogoutRequest]) (TokenResponse, error) {
	req, err := c.Body()
	if err != nil {
		return TokenResponse{}, err
	}

	username := c.Context().Value("username").(string)
	_, err = r.DB.ExecContext(c.Context(), "UPDATE users SET revoke_token_before = NOW() WHERE username = $1", username)
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{Token: req.Token}, nil
}
