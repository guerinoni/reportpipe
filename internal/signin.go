package internal

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SignInHandler struct {
	DB     *sql.DB
	getEnv func(string) string
}

type SignInRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s SignInRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)
	if s.Name == "" {
		problems["name"] = "is required"
	}
	if s.Email == "" {
		problems["email"] = "is required"
	}
	if s.Password == "" {
		problems["password"] = "is required"
	}

	return
}

type SignInResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (s SignInHandler) Path() string {
	return "/signin"
}

func (s SignInHandler) Handler() http.Handler {
	return s
}

func (s SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, problems, err := decodeValid[SignInRequest](r)
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

	_, err = s.DB.ExecContext(r.Context(), "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", req.Name, req.Email, pw)
	if err != nil {
		apiError := ApiError{Errors: map[string]string{"database": err.Error()}}
		if err := encode[ApiError](w, http.StatusInternalServerError, apiError); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	res := SignInResponse{Email: req.Email, Name: req.Name}
	token, err := generateJWT(res.Email, res.Name, s.getEnv)
	if err != nil {
		apiError := ApiError{Errors: map[string]string{"jwt": err.Error()}}
		if err := encode[ApiError](w, http.StatusInternalServerError, apiError); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	res.Token = token
	err = encode[SignInResponse](w, http.StatusCreated, res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

//func checkPassword(providedPassword string, userP) error {
//	err := bcrypt.CompareHashAndPassword([]byte(Password), []byte(providedPassword))
//	if err != nil {
//		return err
//	}
//	return nil
//}
