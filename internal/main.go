// Package generate creates all attribute functions bases on the html standard.
// Probably generation of the elements will also be added.
package main

import (
	_ "embed"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/KrischanCS/htmfunc/internal/attributes"
	"github.com/KrischanCS/htmfunc/internal/elements"
	"github.com/KrischanCS/htmfunc/internal/standard"
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

	start := time.Now()
	defer func() {
		slog.Info("Done", "time", time.Since(start))
	}()

	slog.Info("Begins generation")

	body := standard.LoadStandardForWebDevs(*reloadStandard)

	elements.GenerateElements(body)

	attributes.GenerateAttributes(body)
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
