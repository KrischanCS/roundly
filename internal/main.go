// Package generate creates all attribute functions bases on the html standard.
// Probably generation of the elements will also be added.
package main

import (
	_ "embed"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/KrischanCS/roundly/internal/attributes"
	"github.com/KrischanCS/roundly/internal/elements"
	"github.com/KrischanCS/roundly/internal/standard"
)

//nolint:gochecknoglobals
var reloadStandard = flag.Bool(
	"reloadHtmlStandard",
	false,
	"reload html standard from the web instead of file system",
)

func main() {
	flag.Parse()
	setupLogger()

	slog.Info("Begins generation")

	start := time.Now()
	defer func() {
		slog.Info("Done", "time", time.Since(start))
	}()

	indicesBody := standard.LoadStandardForWebDevsIndices(*reloadStandard)
	syntaxBody := standard.LoadStandardForWebDevsSyntax(*reloadStandard)
	inputBody := standard.LoadStandardForWebDevsInput(*reloadStandard)

	elements.GenerateElements(indicesBody, syntaxBody)
	attributes.GenerateAttributes(indicesBody, inputBody)
}

func setupLogger() {
	slog.SetDefault(
		slog.New(slog.NewTextHandler(os.Stderr,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			}),
		),
	)
}
