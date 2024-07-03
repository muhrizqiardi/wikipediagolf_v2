package dbsetup

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func Migrate(ctx context.Context, databaseURL string) error {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}
	srcDriver, err := iofs.New(postgresMigrationsFS, "postgres_migrations")
	if err != nil {
		return err
	}
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	migration, err := migrate.NewWithInstance(
		"iofs",
		srcDriver,
		"postgres",
		dbDriver,
	)
	if err != nil {
		return err
	}

	slog.Info("migrating all the way to the top")
	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
