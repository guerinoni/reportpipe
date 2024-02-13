package internal

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
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
	DB *sql.DB
}

func (b LoginHandler) Path() string {
	return "/login"
}

func (b LoginHandler) Handler() http.Handler {
	return b
}

type ApiError struct {
	Errors map[string]string `json:"errors"`
}

func (b LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, problems, err := decodeValid[LoginRequest](r)
	if err != nil {
		if err := encode[ApiError](w, http.StatusBadRequest, ApiError{problems}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	rows, err := b.DB.QueryContext(r.Context(), "SELECT * FROM users WHERE email = $1 AND password = $2", req.Email, req.Password)
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

	http.Error(w, "ok", http.StatusOK)
}
