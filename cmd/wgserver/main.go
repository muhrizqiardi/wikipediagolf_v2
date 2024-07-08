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
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	authmiddleware "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/middleware"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signin"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signinpage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signout"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signup"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/signuppage"
	authrepository "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/repository"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/asset"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/config"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/dbsetup"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/home"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/partials"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/feature/gamepage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/feature/gameresult"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/feature/pregame"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/feature/surrender"
	roomcheck "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/check"
	roomcreate "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/create"
	roomcreatepage "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/createpage"
	roomjoin "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/join"
	roomjoinpage "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/joinpage"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/nicknamedialog"
	roomwaitingpage "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/feature/waitingpage"
	featureUsernameCreate "github.com/muhrizqiardi/wikipediagolf_v2/internal/username/feature/create"
	createUsernameModal "github.com/muhrizqiardi/wikipediagolf_v2/internal/username/feature/createmodal"
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
	firebaseApp, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(cfg.FirebaseConfig))
	if err != nil {
		return err
	}
	db, err := dbsetup.Setup(context.Background(), cfg.DatabaseURL, cfg.IsMigrate)
	if err != nil {
		return err
	}
	serveMux := http.NewServeMux()
	tmpl := template.New("")

	actx := authcontext.NewAuthContext()
	authRepository := authrepository.NewRepository(context.Background(), firebaseApp)
	authService := authmiddleware.NewService(authRepository)
	amw := authmiddleware.AuthMiddleware(authService, actx)

	partials.Register(tmpl)
	signuppage.Register(tmpl, serveMux)
	home.Register(tmpl, serveMux, actx)
	asset.Register(serveMux)
	signinpage.Register(tmpl, serveMux)
	roomcreate.Register(context.Background(), db, firebaseApp, serveMux)
	roomcreatepage.Register(tmpl, serveMux)
	roomjoinpage.Register(tmpl, serveMux)
	roomwaitingpage.Register(context.Background(), db, firebaseApp, tmpl, serveMux)
	gamepage.Register(tmpl, serveMux)
	surrender.Register(tmpl, serveMux)
	gameresult.Register(tmpl, serveMux)
	pregame.Register(tmpl, serveMux)
	signin.Register(context.Background(), firebaseApp, serveMux)
	signout.Register(serveMux)
	signup.Register(context.Background(), firebaseApp, tmpl, serveMux)
	createUsernameModal.Register(context.Background(), db, tmpl, serveMux, actx)
	featureUsernameCreate.BuildCreate(context.Background(), db, tmpl, serveMux)
	nicknamedialog.Register(tmpl, serveMux, actx)
	roomcheck.Register(context.Background(), db, firebaseApp, tmpl, serveMux)
	roomjoin.Register(context.Background(), db, firebaseApp, serveMux)

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	slog.Info("starting server", "addr", addr)
	return http.ListenAndServe(addr, amw(serveMux))
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args[1:], os.Getenv, os.Stdin, os.Stdout, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		return
	}
}
