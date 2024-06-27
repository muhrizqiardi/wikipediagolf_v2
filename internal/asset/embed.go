package asset

import "embed"

//go:embed assets/*
var assetFS embed.FS
