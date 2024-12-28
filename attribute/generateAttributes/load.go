package main

import (
	"errors"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
)

func loadIndicesFromStandard() *html.Node {
	if *reloadStandard || !isStandardFilePresent() {
		downloadStandardFile()
	}

	htmlStandard, err := os.Open("htmlStandardIndices.html")
	if err != nil {
		log.Panic("Error opening htmlStandardIndices.html: ", err)
	}

	body, err := html.Parse(htmlStandard)
	if err != nil {
		log.Fatal("Error reading htmlStandardIndices.html: ", err)
	}

	return body
}

func isStandardFilePresent() bool {
	_, err := os.Stat("htmlStandardIndices.html")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		log.Panic("Error checkin if htmlStandardIndices.html exists: ", err)
	}

	return true
}

func downloadStandardFile() {
	const indicesUrl = htmlStandardUrl + "indices.html"

	response, err := http.Get(indicesUrl)
	if err != nil {
		log.Panic("Error loading indices from standard: ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print("Error closing body: ", err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		log.Panic("Unexpected status loading indices from standard: ", response.StatusCode)
	}

	file, err := os.OpenFile("htmlStandardIndices.html", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
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
}
