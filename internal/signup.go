package internal

import (
	"context"
	"database/sql"
	"net/http"

	"reportpipe/internal/auth"

	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s SignUpRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)
	if s.Username == "" {
		problems["username"] = "is required"
	}
	if s.Email == "" {
		problems["email"] = "is required"
	}
	if s.Password == "" {
		problems["password"] = "is required"
	}

	return
}

type SignUpResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type SignUpHandler struct {
	DB     *sql.DB
	getEnv func(string) string
}

func (h SignUpHandler) Path() string {
	return "POST /signup"
}

func (h SignUpHandler) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, problems, err := decodeValid[SignUpRequest](r)
		if err != nil {
			if err := encode[ApiError](w, http.StatusBadRequest, ApiError{problems}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}

		pw, err := hashPassword(req.Password)
		if err != nil {
			apiError := ApiError{Errors: map[string]string{"password": err.Error()}}
			if err := encode[ApiError](w, http.StatusInternalServerError, apiError); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		_, err = h.DB.ExecContext(r.Context(), "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", req.Username, req.Email, pw)
		if err != nil {
			apiError := ApiError{Errors: map[string]string{"database": err.Error()}}
			if err := encode[ApiError](w, http.StatusInternalServerError, apiError); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		res := SignUpResponse{Email: req.Email, Username: req.Username}
		token, err := auth.GenerateJWT(res.Email, res.Username, h.getEnv)
		if err != nil {
			apiError := ApiError{Errors: map[string]string{"jwt": err.Error()}}
			if err := encode[ApiError](w, http.StatusInternalServerError, apiError); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		res.Token = token
		err = encode[SignUpResponse](w, http.StatusCreated, res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func (h SignUpHandler) NeedsAuth() bool {
	return false
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
