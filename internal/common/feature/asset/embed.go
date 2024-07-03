package asset

import "embed"

//go:embed dist/*
var distFS embed.FS

//go:embed assets/*
var assetFS embed.FS
