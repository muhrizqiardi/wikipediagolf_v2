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

	_ "github.com/lib/pq"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/authmiddleware"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/config"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/dbsetup"
	featureSignup "github.com/muhrizqiardi/wikipediagolf_v2/internal/user/feature/signup"
	featureUsernameCreate "github.com/muhrizqiardi/wikipediagolf_v2/internal/username/feature/create"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/username/feature/createpage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/asset"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/game"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/gameresult"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/home"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/view/pregame"
	createroom "github.com/muhrizqiardi/wikipediagolf_v2/internal/view/roomcreate"
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
	db, err := dbsetup.Setup(context.Background(), cfg.DatabaseURL, cfg.IsMigrate)
	if err != nil {
		return err
	}

	serveMux := http.NewServeMux()
	tmpl := template.New("")
	tmpl, err = signup.AddTemplate(tmpl)
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
	firebaseApp, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(cfg.FirebaseConfig))
	if err != nil {
		return err
	}
	signupUserRepository := featureSignup.NewRepository(context.Background(), firebaseApp)
	signupService := featureSignup.NewService(context.Background(), signupUserRepository)
	tmpl, err = featureSignup.AddTemplate(tmpl)
	if err != nil {
		return err
	}
	featureSignup.AddEndpoint(serveMux, featureSignup.EndpointDeps{
		Service:  signupService,
		Template: tmpl,
	})
	featureSignup.Handler(signupService, tmpl)
	featureUsernameCreate.NewRepository(context.Background(), db)
	tmpl, err = featureUsernameCreate.AddTemplate(tmpl)
	if err != nil {
		return err
	}
	featureUsernameCreate.AddEndpoint(serveMux, featureUsernameCreate.EndpointDeps{
		Template: tmpl,
		Service:  nil,
	})
	tmpl, err = createpage.AddTemplate(tmpl)
	if err != nil {
		return err
	}
	createpage.AddEndpoint(serveMux, createpage.EndpointDeps{
		Template: tmpl,
	})

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	slog.Info("starting server", "addr", addr)
	return http.ListenAndServe(addr, authmiddleware.AuthMiddleware(firebaseApp)(serveMux))
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args[1:], os.Getenv, os.Stdin, os.Stdout, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		return
	}
}
