// Package generate creates all attribute functions bases on the html standard.
// Probably generation of the elements will also be added.
package main

import (
	_ "embed"
	"flag"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
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

	indicesBody := standard.LoadStandardForWebDevsIndices(*reloadStandard)
	syntaxBody := standard.LoadStandardForWebDevsSyntax(*reloadStandard)
	inputBody := standard.LoadStandardForWebDevsInput(*reloadStandard)

	elements.GenerateElements(indicesBody, syntaxBody)
	attributes.GenerateAttributes(indicesBody, inputBody)

	copyLogicAndTextFiles()
}

func copyLogicAndTextFiles() {
	filepath.Walk("src/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Panic(err)
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}

		slog.Info("Copying file...", "path", path)

		src, err := os.Open(path)
		if err != nil {
			log.Panic(err)
		}
		defer src.Close()

		dst, err := os.OpenFile("../../html/"+info.Name(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Panic(err)
		}
		defer dst.Close()

		_, err = dst.Write([]byte("// Copied by generator. DO NOT EDIT.\n" +
			"// Original file: " + path + "\n\n"))
		if err != nil {
			log.Panic(err)
		}

		_, err = io.Copy(dst, src)
		if err != nil {
			log.Panic(err)
		}

		slog.Info("Copied file.", "path", path)
		return nil
	})
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
