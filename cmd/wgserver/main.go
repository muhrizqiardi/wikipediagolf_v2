package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/asset"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/config"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/createroompage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/gamepage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/homepage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/joinroompage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/resultpage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/signinpage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/signuppage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/surrenderpage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/waitingroompage"
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
	tmpl := template.New("")
	tmpl, err := signuppage.AddTemplate(tmpl)
	if err != nil {
		return err
	}
	tmpl, err = homepage.AddTemplate(tmpl)
	if err != nil {
		return err
	}
	asset.AddEndpoint(serveMux)
	homepageEndpointDeps := homepage.EndpointDeps{
		Template: tmpl,
	}
	homepage.AddEndpoint(serveMux, homepageEndpointDeps)
	signuppageEndpointDeps := signuppage.EndpointDeps{
		Template: tmpl,
	}
	signuppage.AddEndpoint(serveMux, signuppageEndpointDeps)
	signinpage.AddTemplate(tmpl)
	signinpageEndpointDeps := signinpage.EndpointDeps{
		Template: tmpl,
	}
	signinpage.AddEndpoint(serveMux, signinpageEndpointDeps)
	createroompage.AddTemplate(tmpl)
	createroompageEndpointDeps := createroompage.EndpointDeps{
		Template: tmpl,
	}
	createroompage.AddEndpoint(serveMux, createroompageEndpointDeps)
	joinroompage.AddTemplate(tmpl)
	joinroompageEndpointDeps := joinroompage.EndpointDeps{
		Template: tmpl,
	}
	joinroompage.AddEndpoint(serveMux, joinroompageEndpointDeps)
	waitingroompage.AddTemplate(tmpl)
	waitingroompageEndpointDeps := waitingroompage.EndpointDeps{
		Template: tmpl,
	}
	waitingroompage.AddEndpoint(serveMux, waitingroompageEndpointDeps)
	gamepage.AddTemplate(tmpl)
	gamepageEndpointDeps := gamepage.EndpointDeps{
		Template: tmpl,
	}
	gamepage.AddEndpoint(serveMux, gamepageEndpointDeps)
	surrenderpage.AddTemplate(tmpl)
	surrenderpageEndpointDeps := surrenderpage.EndpointDeps{
		Template: tmpl,
	}
	surrenderpage.AddEndpoint(serveMux, surrenderpageEndpointDeps)
	resultpage.AddTemplate(tmpl)
	resultpageEndpointDeps := resultpage.EndpointDeps{
		Template: tmpl,
	}
	resultpage.AddEndpoint(serveMux, resultpageEndpointDeps)

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
