package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler interface {
	Path() string
	Handler() http.Handler
}

type HealthHandler struct{}

func (h HealthHandler) Path() string {
	return "/health"
}

func (h HealthHandler) Handler() http.Handler {
	return h
}

func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("OK\n"))
}

func AllRoutes(db *sql.DB) []Handler {
	var r []Handler

	r = append(r, HealthHandler{})

	//r = append(r, route{
	//	path:   "/api/v1/admin/count/users",
	//	method: http.MethodGet,
	//	handler: adminOnly(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		w.Write([]byte("42\n"))
	//	})),
	//})
	//
	//r = append(r, route{
	//	path:    "/login",
	//	method:  http.MethodPost,
	//	handler: LoginHandler{DB},
	//})

	r = append(r, LoginHandler{DB: db})

	return r
}

//func adminOnly(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if !currentUser(r).IsAdmin {
//			http.NotFound(w, r)
//			return
//		}
//h.ServeHTTP(w, r)
//})
//}

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func decodeValid[T Validator](r *http.Request) (T, map[string]string, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, nil, fmt.Errorf("decode json: %w", err)
	}
	if problems := v.Valid(r.Context()); len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T: %d problems", v, len(problems))
	}
	return v, nil, nil
}

// openDB opens a connection to a database and returns the connection.
func openDB() (*sql.DB, error) {
	info := "host=localhost user=user password=password dbname=postgres port=5432 sslmode=disable" // TODO: use env or config
	db, err := sql.Open("postgres", info)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}