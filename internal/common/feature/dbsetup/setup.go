package dbsetup

import (
	"context"
	"database/sql"
)

func Setup(ctx context.Context, databaseURL string, isMigrate bool) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	if isMigrate {
		if err := Migrate(ctx, databaseURL); err != nil {
			return nil, err
		}
	}

	return db, err
}
