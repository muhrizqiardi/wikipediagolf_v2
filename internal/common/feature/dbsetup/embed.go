package dbsetup

import "embed"

//go:embed postgres_migrations/*
var postgresMigrationsFS embed.FS
