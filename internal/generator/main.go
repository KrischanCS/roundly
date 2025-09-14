// Package generate creates all attribute functions bases on the html standard.
// Probably generation of the elements will also be added.
package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/KrischanCS/roundly/internal/generator/attributes"
	"github.com/KrischanCS/roundly/internal/generator/elements"
	"github.com/KrischanCS/roundly/internal/generator/standard"
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

	if err := deleteGeneratedFiles(); err != nil {
		slog.Error("failed to delete generated files", "error", err)
		os.Exit(1)
	}

	indicesBody := standard.LoadStandardForWebDevsIndices(*reloadStandard)
	syntaxBody := standard.LoadStandardForWebDevsSyntax(*reloadStandard)
	inputBody := standard.LoadStandardForWebDevsInput(*reloadStandard)

	elements.GenerateElements(indicesBody, syntaxBody)
	attributes.GenerateAttributes(indicesBody, inputBody)
}

func deleteGeneratedFiles() error {
	const header = "// Generated file."

	entries, err := os.ReadDir("../../html")
	if err != nil {
		return fmt.Errorf("read html directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := "../../html/" + entry.Name()

		// Only consider regular files
		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("stat file %s: %w", path, err)
		}
		if !info.Mode().IsRegular() {
			continue
		}

		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("open file %s: %w", path, err)
		}
		// Read the beginning of the file to check the header
		buf := make([]byte, len(header))
		n, err := io.ReadFull(f, buf)
		if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
			_ = f.Close()
			return fmt.Errorf("read header %s: %w", path, err)
		}
		_ = f.Close()

		if string(buf[:n]) == header {
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("delete generated file %s: %w", path, err)
			}
			slog.Info("deleted generated file", "path", path)
		}
	}
	return nil
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
