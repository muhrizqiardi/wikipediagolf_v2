package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/asset"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/homepage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/shared/config"
)

func run(
	_ context.Context,
	args []string,
	getenv func(string) string,
	_ io.Reader,
	stdout,
	_ io.Writer,
) error {
	slog.SetDefault(slog.New(slog.NewTextHandler(stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})))
	cfg := config.GetConfig(args, getenv)

	serveMux := http.NewServeMux()
	tmpl := homepage.MustNewTemplate(homepage.NewTemplate())
	homepageEndpointDeps := homepage.EndpointDeps{
		Template: tmpl,
	}
	homepage.AddEndpoint(serveMux, homepageEndpointDeps)

	return http.ListenAndServe(cfg.Host+":"+strconv.Itoa(cfg.Port), serveMux)
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		return
	}
}
