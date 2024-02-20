package internal

import (
	"context"
	"database/sql"
	"net/http"

	"reportpipe/internal/auth"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l LoginRequest) Valid(_ context.Context) map[string]string {
	problems := make(map[string]string)

	if l.Email == "" {
		problems["email"] = "is required"
	}

	if l.Password == "" {
		problems["password"] = "is required"
	}

	return problems
}

type LoginHandler struct {
	DB     *sql.DB
	getEnv func(string) string
}

func (h LoginHandler) Path() string {
	return "POST /login"
}

func (h LoginHandler) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, problems, err := decodeValid[LoginRequest](r)
		if err != nil {
			if err := encode[ApiError](w, http.StatusBadRequest, ApiError{problems}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}

		rows, err := h.DB.QueryContext(r.Context(), "SELECT username,password FROM users WHERE email = $1", req.Email)
		if err != nil {
			apiError := ApiError{Errors: map[string]string{"database": err.Error()}}
			if err := encode[ApiError](w, http.StatusInternalServerError, apiError); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		defer rows.Close()

		if !rows.Next() {
			apiErr := ApiError{Errors: map[string]string{"user": "not found"}}
			if err := encode[ApiError](w, http.StatusUnauthorized, apiErr); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		var username, pw string
		if err := rows.Scan(&username, &pw); err != nil {
			apiErr := ApiError{Errors: map[string]string{"database": err.Error()}}
			if err := encode[ApiError](w, http.StatusInternalServerError, apiErr); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err := checkPassword(req.Password, pw); err != nil {
			apiErr := ApiError{Errors: map[string]string{"password": "invalid"}}
			if err := encode[ApiError](w, http.StatusUnauthorized, apiErr); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		res := SignUpResponse{Email: req.Email, Username: username}
		token, err := auth.GenerateJWT(res.Email, res.Username, h.getEnv)
		if err != nil {
			apiError := ApiError{Errors: map[string]string{"jwt": err.Error()}}
			if err := encode[ApiError](w, http.StatusInternalServerError, apiError); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		res.Token = token
		err = encode[SignUpResponse](w, http.StatusOK, res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func (h LoginHandler) NeedsAuth() bool {
	return false
}

func checkPassword(providedPassword string, pw string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
