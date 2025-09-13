// Package standard loads the html-standard, either from the stored file or
// refreshes it from the web.
package standard

import (
	"archive/zip"
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
	var filePath = filepath.Join("data", "htmlStandardIndices.zip")
	return loadFile(reload, filePath, "indices.html")
}

func LoadStandardForWebDevsInput(reload bool) *html.Node {
	var filePath = filepath.Join("data", "htmlStandardInputs.zip")
	return loadFile(reload, filePath, "input.html")
}

func LoadStandardForWebDevsSyntax(reload bool) *html.Node {
	var filePath = filepath.Join("data", "htmlStandardSyntax.zip")
	return loadFile(reload, filePath, "syntax.html")
}

func loadFile(reload bool, fileName string, urlResourceName string) *html.Node {
	if reload || !isFilePresent(fileName) {
		downloadStandardIndicesFile(fileName, urlResourceName)
	}

	slog.Info("Parsing HTML file to nodes...", "filePath", fileName)

	//nolint:gosec,nolintlint
	zipReader, err := zip.OpenReader(fileName)
	if err != nil {
		log.Panic("Error opening "+fileName+": ", err)
	}

	file, err := zipReader.Open(urlResourceName)
	if err != nil {
		log.Panic("Error opening "+urlResourceName+": ", err)
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

	saveToFile(fileName, urlResourceName, response.Body)

	slog.Info("Saved HTML standard to file", "file", fileName)
}

func saveToFile(fileName string, urlResourceName string, response io.Reader) {
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

	writeAsZip(file, urlResourceName, response)
}

func writeAsZip(file io.Writer, urlResourceName string, response io.Reader) {
	zipWriter := zip.NewWriter(file)
	defer func() {
		err := zipWriter.Close()
		if err != nil {
			log.Print("Error closing zip writer: ", err)
		}
	}()

	w, err := zipWriter.Create(urlResourceName)
	if err != nil {
		log.Panic("Error creating zip writer for htmlStandardIndices.html: ", err)
	}

	_, err = io.Copy(w, response)
	if err != nil {
		log.Panic("Error saving to htmlStandardIndices.html: ", err)
	}
}
