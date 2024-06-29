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

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/config"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/asset"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/game"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/gameresult"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/home"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/pregame"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/roomcreate"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/roomjoin"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/roomwaiting"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/signin"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/signup"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/surrender"
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
	tmpl, err := signup.AddTemplate(tmpl)
	if err != nil {
		return err
	}
	tmpl, err = home.AddTemplate(tmpl)
	if err != nil {
		return err
	}
	asset.AddEndpoint(serveMux)
	homepageEndpointDeps := home.EndpointDeps{
		Template: tmpl,
	}
	home.AddEndpoint(serveMux, homepageEndpointDeps)
	signuppageEndpointDeps := signup.EndpointDeps{
		Template: tmpl,
	}
	signup.AddEndpoint(serveMux, signuppageEndpointDeps)
	signin.AddTemplate(tmpl)
	signinEndpointDeps := signin.EndpointDeps{
		Template: tmpl,
	}
	signin.AddEndpoint(serveMux, signinEndpointDeps)
	createroom.AddTemplate(tmpl)
	createroompageEndpointDeps := createroom.EndpointDeps{
		Template: tmpl,
	}
	createroom.AddEndpoint(serveMux, createroompageEndpointDeps)
	roomjoin.AddTemplate(tmpl)
	joinroompageEndpointDeps := roomjoin.EndpointDeps{
		Template: tmpl,
	}
	roomjoin.AddEndpoint(serveMux, joinroompageEndpointDeps)
	roomwaiting.AddTemplate(tmpl)
	waitingroompageEndpointDeps := roomwaiting.EndpointDeps{
		Template: tmpl,
	}
	roomwaiting.AddEndpoint(serveMux, waitingroompageEndpointDeps)
	game.AddTemplate(tmpl)
	gamepageEndpointDeps := game.EndpointDeps{
		Template: tmpl,
	}
	game.AddEndpoint(serveMux, gamepageEndpointDeps)
	surrender.AddTemplate(tmpl)
	surrenderpageEndpointDeps := surrender.EndpointDeps{
		Template: tmpl,
	}
	surrender.AddEndpoint(serveMux, surrenderpageEndpointDeps)
	gameresult.AddTemplate(tmpl)
	resultpageEndpointDeps := gameresult.EndpointDeps{
		Template: tmpl,
	}
	gameresult.AddEndpoint(serveMux, resultpageEndpointDeps)
	pregame.AddTemplate(tmpl)
	pregamesplashscreenEndpointDeps := pregame.EndpointDeps{
		Template: tmpl,
	}
	pregame.AddEndpoint(serveMux, pregamesplashscreenEndpointDeps)

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
