package config

import (
	"flag"
	"log/slog"
)

type Config struct {
	Host           string
	Port           int
	IsMigrate      bool
	IsSeed         bool
	DatabaseURL    string
	FirebaseConfig string
}

func GetConfig(args []string, getenv func(string) string) Config {
	f := flag.NewFlagSet("wgserver", flag.PanicOnError)
	var (
		host           = f.String("host", "0.0.0.0", "Listen on host")
		port           = f.Int("port", 3000, "Listen on port")
		isMigrate      = f.Bool("migrate", false, "Migrate database")
		isSeed         = f.Bool("seed", false, "Seed database")
		databaseURL    = getenv("DATABASE_URL")
		firebaseConfig = getenv("FIREBASE_CONFIG")
	)
	f.Parse(args)

	config := Config{
		Host:           *host,
		Port:           *port,
		IsMigrate:      *isMigrate,
		IsSeed:         *isSeed,
		DatabaseURL:    databaseURL,
		FirebaseConfig: firebaseConfig,
	}

	if config.DatabaseURL == "" {
		slog.Warn("environment variable `DATABASE_URL` is empty")
	}

	return config
}
