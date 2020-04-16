package assets

//go:generate go run github.com/gobuffalo/packr/packr

import "github.com/gobuffalo/packr"

var Public = packr.NewBox("./public")
var Docs = packr.NewBox("./docs")
