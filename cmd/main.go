package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	_ "github.com/lib/pq"

	"reportpipe/internal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := internal.Run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
