package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-fuego/fuego"
	"github.com/rs/cors"
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

	server := fuego.NewServer(fuego.WithPort(":8080"))

	fuego.Use(server, cors.New(cors.Options{
		AllowedOrigins: []string{"localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler)

	routes := newRoutes(db, getenv)
	routes.mount(server)

	go func() {
		<-ctx.Done()
		server.Server.Shutdown(ctx)

		stdout.Write([]byte("server shutdown\n"))
	}()

	if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
