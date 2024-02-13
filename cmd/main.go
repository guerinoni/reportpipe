package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
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
