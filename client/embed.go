package client

import "embed"

//go:embed dist/*
var DistFS embed.FS

//go:embed assets/*
var AssetFS embed.FS
