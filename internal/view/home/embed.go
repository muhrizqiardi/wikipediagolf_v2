package home

import "embed"

//go:embed template/*.html
var templateFS embed.FS
