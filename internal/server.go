package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Run(ctx context.Context, args []string, getenv func(string) string, stdin io.Reader, stdout, stderr io.Writer) error {
	db, err := openDB()
	if err != nil {
		msg := fmt.Sprintf("open db: %s\n", err)
		stderr.Write([]byte(msg))
		return fmt.Errorf("open db: %w", err)
	}

	stdout.Write([]byte("db connected\n"))

	_ = db

	mux := http.NewServeMux()
	for _, h := range AllRoutes(db, getenv) {
		mux.Handle(h.Path(), h.Handler())
	}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		server.Shutdown(ctx)

		stdout.Write([]byte("server shutdown\n"))
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
