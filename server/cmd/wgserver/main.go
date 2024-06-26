package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/muhrizqiardi/wikipediagolf_v2/server/cmd/wgserver/internal/httproutes"
)

func run(
	ctx context.Context,
	args []string,
	getenv func(string) string,
	stdin io.Reader,
	stdout,
	stderr io.Writer,
) error {
	handlers := http.NewServeMux()
	httproutes.AddRoutes(handlers)

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		return
	}
}
