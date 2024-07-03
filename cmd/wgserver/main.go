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
	authmiddleware "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/middleware"
	featureSignin "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signin"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signinpage"
	featureSignup "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signup"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signuppage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/asset"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/config"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/dbsetup"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/home"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/feature/game"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/feature/gameresult"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/feature/pregame"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/feature/surrender"
	roomcreatepage "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/createpage"
	roomjoinpage "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/joinpage"
	roomwaitingpage "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/waitingpage"
	featureUsernameCreate "github.com/muhrizqiardi/wikipediagolf_v2/internal/username/feature/create"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/username/feature/createpage"
	usernameMiddleware "github.com/muhrizqiardi/wikipediagolf_v2/internal/username/feature/middleware"
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
	tmpl, err = signuppage.AddTemplate(tmpl)
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
	signuppageEndpointDeps := signuppage.EndpointDeps{
		Template: tmpl,
	}
	signuppage.AddEndpoint(serveMux, signuppageEndpointDeps)
	signinpage.AddTemplate(tmpl)
	signinEndpointDeps := signinpage.EndpointDeps{
		Template: tmpl,
	}
	signinpage.AddEndpoint(serveMux, signinEndpointDeps)
	roomcreatepage.AddTemplate(tmpl)
	createroompageEndpointDeps := roomcreatepage.EndpointDeps{
		Template: tmpl,
	}
	roomcreatepage.AddEndpoint(serveMux, createroompageEndpointDeps)
	roomjoinpage.AddTemplate(tmpl)
	joinroompageEndpointDeps := roomjoinpage.EndpointDeps{
		Template: tmpl,
	}
	roomjoinpage.AddEndpoint(serveMux, joinroompageEndpointDeps)
	roomwaitingpage.AddTemplate(tmpl)
	waitingroompageEndpointDeps := roomwaitingpage.EndpointDeps{
		Template: tmpl,
	}
	roomwaitingpage.AddEndpoint(serveMux, waitingroompageEndpointDeps)
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
	signinRepository := featureSignin.NewRepository(context.Background(), firebaseApp)
	signinService := featureSignin.NewService(signinRepository)
	featureSigninEndpointDeps := featureSignin.EndpointDeps{
		Service: signinService,
	}
	featureSignin.AddEndpoint(serveMux, featureSigninEndpointDeps)
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
	umwr := usernameMiddleware.NewRepository(context.Background(), db)
	umws := usernameMiddleware.NewService(umwr)
	umw := usernameMiddleware.Middleware(umws)

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	slog.Info("starting server", "addr", addr)
	return http.ListenAndServe(addr, umw(authmiddleware.AuthMiddleware(firebaseApp)(serveMux)))
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args[1:], os.Getenv, os.Stdin, os.Stdout, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		return
	}
}
