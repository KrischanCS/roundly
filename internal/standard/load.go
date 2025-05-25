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

//nolint:gochecknoglobals
var standardFileName = filepath.Join("data", "htmlStandardIndices.html")

func LoadStandardForWebDevs(reload bool) *html.Node {
	if reload || !isStandardFilePresent() {
		downloadStandardFile()
	}

	slog.Info("Parsing HTML standard to nodes...", "filePath", standardFileName)

	//nolint:gosec
	htmlStandard, err := os.Open(standardFileName)
	if err != nil {
		log.Panic("Error opening "+standardFileName+": ", err)
	}

	body, err := html.Parse(htmlStandard)
	if err != nil {
		log.Fatal("Error reading "+standardFileName+": ", err)
	}

	slog.Info("Parsed HTML standard to nodes.")

	return body
}

func isStandardFilePresent() bool {
	_, err := os.Stat(standardFileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		log.Panic("Error checkin if htmlStandardIndices.html exists: ", err)
	}

	return true
}

func downloadStandardFile() {
	const indicesURL = HTMLStandardURL + "indices.html"

	slog.Info("Downloading HTML standard from the web...", "url", indicesURL)

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
	file, err := os.OpenFile(standardFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
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

	slog.Info("Saved HTML standard to file", "file", standardFileName)
}
