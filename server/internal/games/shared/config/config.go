package config

import (
	"flag"
	"log/slog"
)

type Config struct {
	Host        string
	Port        int
	IsMigrate   bool
	IsSeed      bool
	DatabaseURL string
}

func GetConfig(args []string, getenv func(string) string) Config {
	f := flag.NewFlagSet("default", flag.PanicOnError)
	var (
		host        = f.String("host", "0.0.0.0", "Listen on host")
		port        = f.Int("port", 3000, "Listen on port")
		isMigrate   = f.Bool("migrate", false, "Migrate database")
		isSeed      = f.Bool("seed", false, "Seed database")
		databaseURL = getenv("DATABASE_URL")
	)
	f.Parse(args)

	config := Config{
		Host:        *host,
		Port:        *port,
		IsMigrate:   *isMigrate,
		IsSeed:      *isSeed,
		DatabaseURL: databaseURL,
	}

	if config.DatabaseURL == "" {
		slog.Warn("environment variable `DATABASE_URL` is empty")
	}

	return config
}
