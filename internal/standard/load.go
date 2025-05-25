// Package standard loads the html-standard, either from the stored file or
// refreshes it from the web.
package standard

import (
	"errors"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

const HTMLStandardURL = `https://html.spec.whatwg.org/dev/`

func LoadStandardForWebDevsIndices(reload bool) *html.Node {
	var filePath = filepath.Join("data", "htmlStandardIndices.html")
	return loadFile(reload, filePath, "indices.html")
}

func LoadStandardForWebDevsInput(reload bool) *html.Node {
	var filePath = filepath.Join("data", "htmlStandardInputs.html")
	return loadFile(reload, filePath, "input.html")
}

func LoadStandardForWebDevsSyntax(reload bool) *html.Node {
	var filePath = filepath.Join("data", "htmlStandardSyntax.html")
	return loadFile(reload, filePath, "syntax.html")
}

func loadFile(reload bool, fileName string, urlResourceName string) *html.Node {
	if reload || !isFilePresent(fileName) {
		downloadStandardIndicesFile(fileName, urlResourceName)
	}

	slog.Info("Parsing HTML file to nodes...", "filePath", fileName)

	//nolint:gosec
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("Error opening "+fileName+": ", err)
	}

	document, err := html.Parse(file)
	if err != nil {
		log.Panic("Error reading "+fileName+": ", err)
	}

	body, ok := findTag(document, "body")
	if !ok {
		log.Panic("Could not find body tag")
	}

	slog.Info("Parsed HTML file to nodes.")

	return body
}

func isFilePresent(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		log.Panicf("Error checking if '%s' exists: %v", filePath, err)
	}

	return true
}

func downloadStandardIndicesFile(fileName string, urlResourceName string) {
	var indicesURL = HTMLStandardURL + urlResourceName

	slog.Info("Downloading HTML standard from the web...", "url", indicesURL)

	//nolint:gosec
	response, err := http.Get(indicesURL)
	if err != nil {
		log.Panic("Error loading indices from standard: ", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Print("Error closing body: ", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		log.Panic("Unexpected status loading indices from standard: ", response.StatusCode)
	}

	//nolint:gosec,mnd,nolintlint
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Print("Error creating file htmlStandardIndices.html: ", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Print("Error closing file: ", err)
		}
	}()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Panic("Error saving to htmlStandardIndices.html: ", err)
	}

	slog.Info("Saved HTML standard to file", "file", fileName)
}
