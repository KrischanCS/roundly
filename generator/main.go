// Package generate creates all attribute functions bases on the html standard.
// Probably generation of the elements will also be added.
package main

import (
	_ "embed"
	"flag"

	"github.com/KrischanCS/htmfunc/generator/internal/attributes"
	"github.com/KrischanCS/htmfunc/generator/internal/elements"
	"github.com/KrischanCS/htmfunc/generator/internal/standard"
)

//nolint:gochecknoglobals
var reloadStandard = flag.Bool("reloadHtmlStandard", false, "reload html standard from the web instead of file system")

func main() {
	flag.Parse()

	body := standard.LoadStandardForWebDevs(*reloadStandard)

	elements.GenerateElements(body)
	attributes.GenerateAttributes(body)
}
