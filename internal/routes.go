package internal

import (
	"database/sql"
	"github.com/go-fuego/fuego"
)

type Routes struct {
	DB     *sql.DB
	getEnv func(string) string
}

func newRoutes(db *sql.DB, getEnv func(string) string) Routes {
	return Routes{
		DB:     db,
		getEnv: getEnv,
	}
}

func (r *Routes) mount(server *fuego.Server) {
	fuego.Get(server, "/health", func(c fuego.ContextNoBody) (string, error) {
		return "OK", nil
	})
	fuego.Post(server, "/login", r.login)
	fuego.Post(server, "/signup", r.signup)
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
