package gongdoc

import "embed"

//go:embed go/models go/diagrams
var GoDir embed.FS
